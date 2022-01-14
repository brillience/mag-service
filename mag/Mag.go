package mag

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/brillience/es_service/elastic"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Abstract struct {
	Id      string `json:"docid"`
	Content string `json:"abstract"`
}

func (a *Abstract) String() string {
	return fmt.Sprintf("Abstract: id->{%s}, content->{%s};", a.Id, a.Content)
}

// loadCsv 将csv文件的摘要数据映射到Abstract
func loadCsv(csvPath string) chan Abstract {
	//创建一个大小为128的带缓冲通道
	session := make(chan Abstract, 128)
	if path.Ext(csvPath) != ".csv" {
		log.Fatalln("You should use .csv format file!")
	}
	file, err := os.Open(csvPath)
	if err != nil {
		log.Fatalln("open csv file err: ", err.Error())
	}
	reader := csv.NewReader(file)
	//采用生成器模式
	go func() {
		//流式读入
		for {
			line, err := reader.Read()
			if err == io.EOF {
				log.Printf("file: %s read over! \n", csvPath)
				break
			}
			if err != nil {
				log.Fatalln("csv read error ：", err.Error())
			}
			session <- Abstract{Id: line[0], Content: line[1]}
		}
		close(session)
	}()
	return session
}

// UpdateCsvToEs 将CSV文件的摘要数据增量式更新到elasticsearch的mag索引上
// @params:path csv文件的路径
func UpdateCsvToEs(csvPath string, client *elastic.Client) {
	// 确认mag的索引是否创建
	client.CheckIndexOrCreat("mag", `
	{
		"mappings": {
			"properties": {
				"docid": {
					"type": "keyword"
				},
				"abstract": {
					"type": "text"
				}
			}
		}
	}
	`)
	// 覆盖性更新文档；若文档不存在则创建
	abstracts := loadCsv(csvPath)
	for {
		abs := <-abstracts
		if abs.Id == "" {
			break
		}
		document := Abstract{Id: abs.Id, Content: abs.Content}
		CreateMagDocument(document, client)
	}
}

// CreateMagDocument 覆盖性更新文档；若文档ID不存在，则创建；
func CreateMagDocument(abs Abstract, client *elastic.Client) {
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(&abs)
	response, err := client.Index("mag", body, client.Index.WithDocumentID(abs.Id))
	defer response.Body.Close()
	if err != nil {
		log.Fatalln("[ERROR] ", err.Error())
	}
	log.Println(response)
}

// UpdateMagDocument 更新文档
func UpdateMagDocument(abs Abstract, client *elastic.Client) {
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(&abs)
	if err != nil {
		log.Fatalln("[ERROR] ", err.Error())
	}
	_, err = client.Update("mag", abs.Id, body)
	if err != nil {
		log.Fatalln("[ERROR] ", err.Error())
	}
}

// GetMagDocumentById 通过摘要的id获取查询结果
func GetMagDocumentById(id string, client *elastic.Client) Abstract {
	response, err := client.Get("mag", id)
	if err != nil {
		log.Fatalln("[ERROR]: ", err.Error())
	}
	defer response.Body.Close()
	respItem := GetReponse{}
	j, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(j, &respItem)
	return Abstract{Id: respItem.Source.Id, Content: respItem.Source.Content}
}

// SearchByKeyWord 支持通配符的模糊搜索
func SearchMagDocumentsByKeyWord(KeyWord string, client *elastic.Client) []Abstract {

	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"wildcard": map[string]interface{}{
				"abstract": KeyWord,
			},
		},
		"size": 5000,
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex("mag"),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	articles := []Abstract{}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		id := hit.(map[string]interface{})["_source"].(map[string]interface{})["docid"].(string)
		content := hit.(map[string]interface{})["_source"].(map[string]interface{})["abstract"].(string)
		articles = append(articles, Abstract{Id: id, Content: content})
	}
	return articles
}
