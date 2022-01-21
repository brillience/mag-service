package svc

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"mag_service/rpc/elastic"
	"mag_service/rpc/internal/config"
	"mag_service/rpc/model"
)

type ServiceContext struct {
	Config       config.Config
	NlpTagsModel model.NlpTagsModel
	MagEs        *elastic.AbstractEs
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{Addresses: c.EsConfig.Urls, Username: c.EsConfig.User, Password: c.EsConfig.Password})
	if err != nil {
		panic(err)
	}
	magEs, err := elastic.NewMagEs(esClient)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:       c,
		NlpTagsModel: model.NewNlpTagsModel(conn, c.CacheRedis),
		MagEs:        magEs,
	}
}
