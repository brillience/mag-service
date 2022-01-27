package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"mag_service/rpc/internal/config"
	"mag_service/rpc/internal/server"
	"mag_service/rpc/internal/svc"
	"mag_service/rpc/mag"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/mag.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewMagServer(ctx)
	// 增量式更新 摘要数据到 ES
	logx.Info("开始更新uploadData/abstracts.csv到elasticsearch...")
	ctx.MagEs.UpdateCsvToEs("./uploadData/abstracts.csv")
	logx.Info("同步更新完毕！")
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		mag.RegisterMagServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
