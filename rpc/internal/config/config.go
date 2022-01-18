package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type ElasticConfig struct {
	Urls     []string
	User     string
	Password string
}
type Config struct {
	zrpc.RpcServerConf
	EsConfig ElasticConfig
	Mysql    struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
}
