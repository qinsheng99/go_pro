package controller

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/qinsheng99/go-domain-web/app"
	"github.com/qinsheng99/go-domain-web/domain/kubernetes"
	"github.com/qinsheng99/go-domain-web/utils"
)

type BaseConfigMap struct {
	c app.ConfigMapServiceImpl
}

func AddRouteConfigMap(r *gin.RouterGroup, c kubernetes.ConfigMap) {
	baseConfigMap := &BaseConfigMap{
		c: app.NewConfigMapService(c),
	}

	group := r.Group("/config-map")

	func() {
		group.POST("/create", baseConfigMap.Create)
	}()
}

func (b *BaseConfigMap) Create(c *gin.Context) {
	err := b.c.Create(context.TODO())
	if err != nil {
		utils.Failure(c, err)
		return
	}

	utils.SuccessCreate(c, "success")
}
