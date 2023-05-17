package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/mysql"
	"github.com/qinsheng99/go-domain-web/common/infrastructure/postgres"
	"github.com/qinsheng99/go-domain-web/common/logger"
	"github.com/qinsheng99/go-domain-web/config"
	"github.com/qinsheng99/go-domain-web/server"
	"github.com/qinsheng99/go-domain-web/task"
	"github.com/qinsheng99/go-domain-web/utils"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	flag.Parse()

	//gin.SetMode(gin.ReleaseMode)
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

	if err = logger.InitLogger(cfg.Logger); err != nil {
		logrus.WithError(err).Error("logger init failed")

		return
	}
	//
	//if err = kubernetes.Init(cfg.Kubernetes); err != nil {
	//	logrus.WithError(err).Error("kubernetes init failed")
	//
	//	return
	//}

	if err = mysql.Init(cfg.Mysql); err != nil {
		logrus.WithError(err).Error("mysql init failed")

		return
	}

	if err = postgres.Init(cfg.Postgres); err != nil {
		logrus.WithError(err).Error("postgres init failed")

		return
	}

	//if err = elastic.Init(cfg.Es); err != nil {
	//	logrus.WithError(err).Error("elasticsearch init failed")
	//
	//	return
	//}
	//
	//if err = redis.Init(cfg.Redis); err != nil {
	//	logrus.WithError(err).Error("redis init failed")
	//
	//	return
	//}
	//
	//if err = mongodb.Init(cfg.Mongo); err != nil {
	//	logrus.WithError(err).Error("mongodb init failed")
	//
	//	return
	//}
	//
	//if err = etcd.Init(cfg.Etcd); err != nil {
	//	logrus.WithError(err).Error("etcd init failed")
	//
	//	return
	//}

	server.SetRoute(r, cfg)

	t := task.NewTask(cfg.Task, cfg.Postgres)
	//if err = t.Register(); err != nil {
	//	logrus.WithError(err).Error("register task failed")
	//
	//	return
	//}

	go t.Pkg()

	//t.Run()
	//defer t.Stop()

	//lis := kubernetes.NewListen(kubernetes.GetClient(), kubernetes.GetDyna(), kubernetes.GetResource(), *listen)
	//go lis.ListenResource()

	utils.Start(cfg.Port, r.Handler())
}
