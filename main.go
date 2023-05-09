package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/mysql"
	"github.com/qinsheng99/go-domain-web/common/infrastructure/postgresql"
	"github.com/qinsheng99/go-domain-web/common/logger"
	"github.com/qinsheng99/go-domain-web/config"
	"github.com/qinsheng99/go-domain-web/server"
	"github.com/qinsheng99/go-domain-web/utils"
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
		logrus.WithError(err).Error("config init failed")

		return
	}

	err = logger.InitLogger(cfg.Logger)
	if err != nil {
		logrus.WithError(err).Error("logger init failed")

		return
	}

	//err = kubernetes.Init(cfg.KubernetesConfig)
	//if err != nil {
	//	logrus.WithError(err).Error("kubernetes init failed")
	//
	//  return
	//}
	err = mysql.Init(cfg.Mysql)
	if err != nil {
		logrus.WithError(err).Error("mysql init failed")

		return
	}

	err = postgresql.Init(cfg.Postgresql)
	if err != nil {
		logrus.WithError(err).Error("postgresql init failed")

		return
	}

	//err = elastic.Init(cfg.Es)
	//if err != nil {
	//	logrus.WithError(err).Error("elasticsearch init failed")
	//
	//	return
	//}
	//
	//err = redis.Init(cfg.Redis)
	//if err != nil {
	//	logrus.WithError(err).Error("redis init failed")
	//
	//	return
	//}
	//
	//err = mongodb.Init(cfg.Mongo)
	//if err != nil {
	//	logrus.WithError(err).Error("mongodb init failed")
	//
	//	return
	//}
	//
	//err = etcd.Init(cfg.Etcd)
	//if err != nil {
	//	logrus.WithError(err).Error("etcd init failed")
	//
	// return
	//}

	//task.RepoTask()
	server.SetRoute(r, cfg)

	//lis := kubernetes.NewListen(kubernetes.GetClient(), kubernetes.GetDyna(), kubernetes.GetResource(), *listen)
	//go lis.ListenResource()

	utils.Start(cfg.Port, r.Handler())
}
