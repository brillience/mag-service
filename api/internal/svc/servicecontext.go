package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
	"mag_service/api/internal/config"
	"mag_service/api/model"
	"mag_service/rpc/magclient"
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
