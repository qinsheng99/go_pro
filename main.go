package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-domain-web/config"
	"github.com/qinsheng99/go-domain-web/infrastructure/kubernetes"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
	"github.com/qinsheng99/go-domain-web/infrastructure/postgresql"
	"github.com/qinsheng99/go-domain-web/logger"
	"github.com/qinsheng99/go-domain-web/route"
	"log"
)

func main() {
	r := gin.Default()

	err := config.Init()
	if err != nil {
		panic(err)
	}

	err = logger.InitLogger(config.Conf.LogConfig)
	if err != nil {
		panic(err)
	}

	err = kubernetes.Init()
	if err != nil {
		panic(err)
	}

	err = mysql.Init(config.Conf.MysqlConfig)
	if err != nil {
		panic(err)
	}

	err = postgresql.Init(config.Conf.PostgresqlConfig)
	if err != nil {
		panic(err)
	}

	//err = elasticsearch.Init(config.Conf.EsConfig)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = redis.Init(config.Conf.RedisConfig)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = mongodb.Init(config.Conf.MongoConfig)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = etcd.Init(config.Conf.EtcdConfig)
	//if err != nil {
	//	panic(err)
	//}

	route.SetRoute(r)

	err = r.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
