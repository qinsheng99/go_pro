package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-py/route"
	"log"
)

func main() {
	r := gin.Default()

	route.SetRoute(r)

	err := r.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
