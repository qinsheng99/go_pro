package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/mysql"
	"github.com/qinsheng99/go-domain-web/common/infrastructure/postgresql"
	"github.com/qinsheng99/go-domain-web/common/logger"
	"github.com/qinsheng99/go-domain-web/config"
	"github.com/qinsheng99/go-domain-web/route"
	"github.com/qinsheng99/go-domain-web/utils/server"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	flag.Parse()
}

var listen = flag.Bool("listen", false, "")
var path = flag.String("config-file", "config/config.yaml", "")

func main() {
	r := gin.Default()

	cfg, err := config.Init(*path)
	if err != nil {
		logrus.WithError(err).Fatal("config init failed")
	}

	err = logger.InitLogger(cfg.Logger)
	if err != nil {
		logrus.WithError(err).Fatal("logger init failed")
	}

	//err = kubernetes.Init(cfg.KubernetesConfig)
	//if err != nil {
	//	logrus.WithError(err).Fatal("kubernetes init failed")
	//}
	err = mysql.Init(cfg.Mysql)
	if err != nil {
		logrus.WithError(err).Fatal("mysql init failed")
	}

	err = postgresql.Init(cfg.Postgresql)
	if err != nil {
		logrus.WithError(err).Fatal("postgresql init failed")
	}

	//err = elastic.Init(cfg.Es)
	//if err != nil {
	//	logrus.WithError(err).Fatal("elasticsearch init failed")
	//}
	//
	//err = redis.Init(cfg.Redis)
	//if err != nil {
	//	logrus.WithError(err).Fatal("redis init failed")
	//}
	//
	//err = mongodb.Init(cfg.Mongo)
	//if err != nil {
	//	logrus.WithError(err).Fatal("mongodb init failed")
	//}
	//
	//err = etcd.Init(cfg.Etcd)
	//if err != nil {
	//	logrus.WithError(err).Fatal("etcd init failed")
	//}

	//task.RepoTask()
	route.SetRoute(r, cfg)

	//lis := kubernetes.NewListen(kubernetes.GetClient(), kubernetes.GetDyna(), kubernetes.GetResource(), *listen)
	//go lis.ListenResource()

	server.Start(cfg.Port, r.Handler())
}
