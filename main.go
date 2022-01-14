package main

import (
	"fmt"
	"github.com/brillience/es_service/elastic"
	"github.com/brillience/es_service/mag"
)

func main() {
	client := elastic.NewClient()
	abstracts := mag.SearchMagDocumentsByKeyWord("stromatoli*", client)
	fmt.Println(len(abstracts))
}
