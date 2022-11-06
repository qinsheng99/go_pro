package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	app "github.com/qinsheng99/go-domain-web/app/kubernetes"
	"github.com/qinsheng99/go-domain-web/domain/kubernetes"
	"github.com/qinsheng99/go-domain-web/utils"
	"net/http"
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
	pod, err := b.p.GetPod(context.TODO(), c.Param("name"))
	if err != nil {
		utils.Failure(c, err)
		return
	}

	utils.Success(c, http.StatusOK, pod)
}

func (b *BasePod) List(c *gin.Context) {
	pod, err := b.p.PodList(context.TODO())
	if err != nil {
		utils.Failure(c, err)
		return
	}

	utils.Success(c, http.StatusOK, pod)
}

func (b *BasePod) Create(c *gin.Context) {
	err := b.p.Create(context.TODO())
	if err != nil {
		utils.Failure(c, err)
		return
	}
	utils.Success(c, http.StatusCreated, "success")
}
