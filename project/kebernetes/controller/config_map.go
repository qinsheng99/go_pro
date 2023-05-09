package controller

import (
	"context"

	"github.com/gin-gonic/gin"

	commonctl "github.com/qinsheng99/go-domain-web/common/controller"
	"github.com/qinsheng99/go-domain-web/project/kebernetes/app"
	"github.com/qinsheng99/go-domain-web/project/kebernetes/domain/kubernetes"
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
	if err := b.c.Create(context.TODO()); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.SuccessCreate(c)
	}
}
