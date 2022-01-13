package elastic

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/jinzhu/configor"
	"log"
	"strings"
)

var Conf = struct {
	Addresses []string `required:"true"`
	Username  string
	Password  string
}{}
var esConfig elasticsearch.Config

func init() {
	//load config
	err := configor.Load(&Conf, "config.yaml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	esConfig.Addresses = Conf.Addresses
	if Conf.Username != "" && Conf.Password != "" {
		esConfig.Username = Conf.Username
		esConfig.Password = Conf.Password
	}

}

type Client struct {
	*elasticsearch.Client
}

// NewClient 创建客户端
func NewClient() *Client {
	client, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return &Client{client}
}

// CheckIndexOrCreat 判断索引是否存在，不存在则创建
func (client Client) CheckIndexOrCreat(indexName string, body string) {
	//判断索引是否存在，不存在则创建
	resp, err := client.Indices.Exists([]string{indexName})
	if err != nil {
		fmt.Println("ERROR!", err.Error())
	}
	//resp的状态为200，则存在该索引；404则不存在该索引
	if resp.StatusCode == 404 {
		response, err := client.Indices.Create("mag", client.Indices.Create.WithBody((strings.NewReader(body))))
		if err != nil {
			log.Fatalf("Cannot create index: %s", err.Error())
		}
		if response.IsError() {
			log.Fatalf("Cannot create index: %s", response)
		}
	}
}
