package main

import (
	"bytes"
	"encoding/json"
	"github.com/brillience/es_service/elastic"
	"github.com/brillience/es_service/mag"
	"log"
)

func main() {
	client := elastic.NewClient()
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
	abstracts := mag.LoadCsv("./articles.csv")
	for {
		abs := <-abstracts
		if abs.Id == "" {
			break
		}
		document := mag.Abstract{Id: abs.Id, Content: abs.Content}
		body := &bytes.Buffer{}
		err := json.NewEncoder(body).Encode(&document)
		response, err := client.Index("mag", body, client.Index.WithDocumentID(document.Id))
		if err != nil {
			log.Fatalln("[ERROR] ", err.Error())
		}
		log.Println(response)
	}
}
