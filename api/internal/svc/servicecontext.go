package svc

import (
	"es_service/api/internal/config"
	"es_service/api/model"
	"es_service/rpc/magclient"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	MagRpc    magclient.Mag
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
		MagRpc:    magclient.NewMag(zrpc.MustNewClient(c.MagRpc)),
	}
}
