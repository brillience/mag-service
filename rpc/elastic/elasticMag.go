package elastic

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/tal-tech/go-zero/core/logx"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

//
//  AbstractEs
//  @Description: 携带elasticsearch客户端的Mag摘要Es句柄
//
type AbstractEs struct {
	index   string
	mapping string
	client  *elasticsearch.Client
}

//
//  Abstract
//  @Description: 用于解析检索结果的json数据
//
type Abstract struct {
	Id      string `json:"docid"`
	Content string `json:"abstract"`
}

//
//  doc
//  @Description: 用于elasticsearch更新数据的结构映射
//
type doc struct {
	Doc interface{} `json:"doc"`
}

//
//  GetReponse
//  @Description: 用于解析elasticsearch的mag索引的单条查询结果的json数据
//
type GetReponse struct {
	Index       string   `json:"_index"`
	Type        string   `json:"_type"`
	Id          string   `json:"_id"`
	Version     int      `json:"_version"`
	SeqNo       int      `json:"_seq_no"`
	PrimaryTerm int      `json:"_primary_term"`
	Found       bool     `json:"found"`
	Source      Abstract `json:"_source"`
}

//
//  NewMagEs
//  @Description: 创建AbstractEs，摘要数据的client
//  @param client
//  @return *AbstractEs
//  @return error
//
func NewMagEs(client *elasticsearch.Client) (*AbstractEs, error) {
	abstractEs := &AbstractEs{
		index: "mag",
		mapping: `
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
	`,
		client: client,
	}
	//判断索引是否存在，不存在则创建
	resp, err := abstractEs.client.Indices.Exists([]string{abstractEs.index})
	if err != nil {
		return nil, err
	}
	//resp的状态为200，则存在该索引；404则不存在该索引
	if resp.StatusCode == 404 {
		_, err := abstractEs.client.Indices.Create(abstractEs.index, client.Indices.Create.WithBody((strings.NewReader(abstractEs.mapping))))
		if err != nil {
			return nil, err
		}
	}
	return abstractEs, nil

}

//
//  LoadCsv
//  @Description: 加载mag的摘要的csv文件返回摘要结构体
//  @receiver abstractEs
//  @param csvPath
//  @return chan
//
func (abstractEs *AbstractEs) loadCsv(csvPath string) chan Abstract {
	//创建一个大小为128的带缓冲通道
	session := make(chan Abstract, 128)
	if path.Ext(csvPath) != ".csv" {
		logx.Error("You should use .csv format file!")
	}
	file, err := os.Open(csvPath)
	if err != nil {
		logx.Error("open csv file err: ", err.Error())
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

//
//  UpdateCsvToEs
//  @Description: 将CSV文件的摘要数据增量式更新到elasticsearch的mag索引上
//  @receiver abstractEs
//  @param csvPath
//
func (abstractEs *AbstractEs) UpdateCsvToEs(csvPath string) {
	// 覆盖性更新文档；若文档不存在则创建
	abstracts := abstractEs.loadCsv(csvPath)
	for {
		abs := <-abstracts
		if abs.Id == "" {
			break
		}
		document := Abstract{Id: abs.Id, Content: abs.Content}
		err := abstractEs.CreateDocument(document)
		if err != nil {
			logx.Error(err.Error())
		}
	}
}

//
//  CreateDocument
//  @Description: 覆盖性更新文档；若文档ID不存在，则创建；
//  @receiver abstractEs
//  @param abs
//  @return error
//
func (abstractEs *AbstractEs) CreateDocument(abs Abstract) error {
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(&abs)
	if err != nil {
		return err
	}
	resp, err := abstractEs.client.Index("mag", body, abstractEs.client.Index.WithDocumentID(abs.Id))
	if resp.IsError() {
		return errors.New(resp.String())
	}
	return err
}

//
//  DeleteDocument
//  @Description: 通过docid删除文档
//  @receiver abstractEs
//  @param docid
//  @return error
//
func (abstractEs *AbstractEs) DeleteDocument(docid string) error {
	response, err := abstractEs.client.Delete(abstractEs.index, docid)
	if response.IsError() {
		return errors.New(response.String())
	}
	return err
}

//
//  UpdateDocument
//  @Description: 更新文档
//  @receiver abstractEs
//  @param abs
//
func (abstractEs *AbstractEs) UpdateDocument(abs Abstract) error {
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(&doc{
		Doc: &abs,
	})
	if err != nil {
		return err
	}
	resp, err := abstractEs.client.Update(abstractEs.index, abs.Id, body)
	if resp.IsError() {
		return errors.New(resp.String())
	}
	return err
}

//
//  GetDocumentById
//  @Description: 通过id查询摘要结果
//  @receiver abstractEs
//  @param id
//  @return *Abstract
//  @return error
//
func (abstractEs *AbstractEs) GetDocumentById(id string) (*Abstract, error) {
	response, err := abstractEs.client.Get("mag", id)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}
	defer response.Body.Close()
	respItem := GetReponse{}
	j, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(j, &respItem)
	if err != nil {
		return nil, err
	}
	return &Abstract{Id: respItem.Source.Id, Content: respItem.Source.Content}, nil
}

//
//  SearchDocumentsByKeyWord
//  @Description: 支持通配符的模糊搜索
//  @receiver abstractEs
//  @param KeyWord
//  @return []Abstract
//  @return error
//
func (abstractEs *AbstractEs) SearchDocumentsByKeyWord(KeyWord string) ([]Abstract, error) {

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
		return nil, err
	}
	res, err := abstractEs.client.Search(
		abstractEs.client.Search.WithContext(context.Background()),
		abstractEs.client.Search.WithIndex("mag"),
		abstractEs.client.Search.WithBody(&buf),
		abstractEs.client.Search.WithTrackTotalHits(true),
		abstractEs.client.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, errors.New(fmt.Sprintf("Error parsing the response body: %s", err))
		} else {
			// Print the response status and error information.
			return nil, errors.New(fmt.Sprintf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			))
		}
	}
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, errors.New(fmt.Sprintf("Error parsing the response body: %s", err))
	}
	// Print the response status, number of results, and request duration.
	logx.Infof(
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
	return articles, nil
}
