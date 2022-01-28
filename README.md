# Mag Service

MAG abstract service for ODD. This service provides both API service and Rpc service.

## Deployment
OS: `Ubuntu >=18.04`
1. Install docker and docker-compose
    ```shell
    sudo apt-get install docker.io
    sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose
    sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
    sudo apt install python3-pip
    ```
2. Put `abstracts.csv` to `./rpc/uploadData`
    ```shell
    mv abstracts.csv ./rpc/uploadData/
    ```
   
3. Mark this `abstracts.csv` with [NLPMarkTool](https://github.com/brillience/NLPMarkTool).
   - Convert csv to tsv (optional)
   ```shell
   curl -L https://github.com/eBay/tsv-utils/releases/download/v2.2.0/tsv-utils-v2.2.0_linux-x86_64_ldc2.tar.gz | tar xz
   ./tsv-utils-v2.2.0_linux-x86_64_ldc2/bin/csv2tsv abstracts.csv > articles.tsv
   ```
   - NLPMark
   ```shell
   git clone https://github.com/brillience/NLPMarkTool
   mv articles.tsv ./NLPMarkTool/ && cd ./NLPMarkTool
   mvn clean
   mvn install
   mvn exec:java -Dexec.mainClass="com.zhang.nlp.Main"
   ```
   Then, `sentences.tsv` file wil be get and put this file at project path:`mag-service`.
4. Build env.
   ```shell
   docker-compose -f docker-compose-env.yaml up -d
   ```
5. Database Init.
   - Create database and tables.
   ```shell
   bash ./scripts/db_init.sh
   ```
   - update nlptags to db.
   ```shell
   bash ./scripts/update_db.sh
   ```
   - update `abstracts.csv` to elasticsearch.
   ```shell
   go run updateEs.go --p=abstracts.csv
   ```
7. Build Server.
   
   ```shell
   docker-compose up -d
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
