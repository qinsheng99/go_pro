package controller

import (
	"context"

	"github.com/gin-gonic/gin"

	commonctl "github.com/qinsheng99/go-domain-web/common/controller"
	"github.com/qinsheng99/go-domain-web/project/kebernetes/app"
	"github.com/qinsheng99/go-domain-web/project/kebernetes/domain/kubernetes"
)

type BasePod struct {
	p app.PodServiceImpl
}

func AddRoutePod(r *gin.RouterGroup, p kubernetes.Pod) {
	basePod := &BasePod{
		p: app.NewPodService(p),
	}

	group := r.Group("/pod")

	func() {
		group.GET("/get-pod/:name", basePod.Get)
		group.GET("/pod-list", basePod.List)
		group.POST("/create", basePod.Create)

	}()
}

func (b *BasePod) Get(c *gin.Context) {
	if pod, err := b.p.GetPod(context.TODO(), c.Param("name")); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.Success(c, pod)
	}
}

func (b *BasePod) List(c *gin.Context) {
	if pod, err := b.p.PodList(context.TODO()); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.Success(c, pod)
	}
}

func (b *BasePod) Create(c *gin.Context) {
	if err := b.p.Create(context.TODO()); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.SuccessCreate(c)
	}
}
