# Mag Service

MAG abstract service for ODD. This service provides both API service and Rpc service.

## Deployment

To deploy this project run

```bash
  xxx
```


## RESTful API 
[API doc.](api/mag.md)

## RPC Demo
Firstly, you should get this project locally. Then how to use it can be found [here](api/internal/svc/servicecontext.go).
And this is another example：
```go
package main

import (
	"flag"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"

	"mag_service/rpc/magclient"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
	MagRpc zrpc.RpcClientConf
}

var configFile = flag.String("f", "api/etc/mag-api.yaml", "the config file")

func main() {
	flag.Parse()
	var c Config
	conf.MustLoad(*configFile, &c)
	client := magclient.NewMag(zrpc.MustNewClient(c.MagRpc))
	client.GetNlpById(...)
}

```

## Authors

- [@brillience](https://github.com/brillience)
