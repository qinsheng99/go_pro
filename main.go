package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-domain-web/config"
	"github.com/qinsheng99/go-domain-web/infrastructure/elasticsearch"
	"github.com/qinsheng99/go-domain-web/infrastructure/kubernetes"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
	"github.com/qinsheng99/go-domain-web/logger"
	"github.com/qinsheng99/go-domain-web/route"
	"github.com/qinsheng99/go-domain-web/utils/server"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	flag.Parse()
}

var listen = flag.Bool("listen", false, "")

func main() {
	r := gin.Default()

	err := config.Init()
	if err != nil {
		logrus.WithError(err).Fatal("config init failed")
	}

	err = logger.InitLogger(config.Conf.LogConfig)
	if err != nil {
		logrus.WithError(err).Fatal("logger init failed")
	}

	//err = kubernetes.Init(config.Conf.KubernetesConfig)
	//if err != nil {
	//	logrus.WithError(err).Fatal("kubernetes init failed")
	//}
	err = mysql.Init(config.Conf.MysqlConfig)
	if err != nil {
		logrus.WithError(err).Fatal("mysql init failed")
	}

	//err = postgresql.Init(config.Conf.PostgresqlConfig)
	//if err != nil {
	//	logrus.WithError(err).Fatal("postgresql init failed")
	//}

	err = elasticsearch.Init(config.Conf.EsConfig)
	if err != nil {
		logrus.WithError(err).Fatal("elasticsearch init failed")
	}
	//
	//err = redis.Init(config.Conf.RedisConfig)
	//if err != nil {
	//	logrus.WithError(err).Fatal("redis init failed")
	//}
	//
	//err = mongodb.Init(config.Conf.MongoConfig)
	//if err != nil {
	//	logrus.WithError(err).Fatal("mongodb init failed")
	//}
	//
	//err = etcd.Init(config.Conf.EtcdConfig)
	//if err != nil {
	//	logrus.WithError(err).Fatal("etcd init failed")
	//}

	route.SetRoute(r)

	lis := kubernetes.NewListen(kubernetes.GetClient(), kubernetes.GetDyna(), kubernetes.GetResource(), *listen)
	go lis.ListenResource()

	server.Start(config.Conf.Port, r.Handler())
}
