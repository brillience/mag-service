package main

import (
	"flag"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"mag_service/rpc/elastic"
	"path"
)

var p string

func init() {
	flag.StringVar(&p, "path", "", "abstracts.csv path")
}
func main() {
	flag.Parse()
	if p == "" || path.Ext(p) != ".csv" {
		panic("Please input csv path")
	}
	client, err := elasticsearch.NewClient(elasticsearch.Config{Addresses: []string{"127.0.0.1:92009"}})
	if err != nil {
		panic(err.Error())
	}
	magEs, err := elastic.NewMagEs(client)
	magEs.UpdateCsvToEs(p)
	fmt.Println("Update to ES done！")
}
