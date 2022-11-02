package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-domain-web/config"
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

	route.SetRoute(r)

	err = r.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
