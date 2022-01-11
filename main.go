package main

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/jinzhu/configor"
	"log"
)

var esConfig elasticsearch.Config
var Conf = struct {
	Addresses []string `required:"true"`
	Username  string
	Password  string
}{}

func init() {
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

func main() {
	es, _ := elasticsearch.NewClient(esConfig)
	log.Println(elasticsearch.Version)
	log.Println(es.Info())
}
