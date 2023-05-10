package controller

import (
	"github.com/gin-gonic/gin"

	commonctl "github.com/qinsheng99/go-domain-web/common/controller"
	"github.com/qinsheng99/go-domain-web/project/cve/app"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
)

type CveController struct {
	repo app.CveService
}

func AddRouteCve(r *gin.RouterGroup, repo app.CveService) {
	ctl := &CveController{
		repo: repo,
	}

	group := r.Group("/cve")

	func() {
		group.POST("/upload", ctl.Upload)

		group.GET("/:cveNum", ctl.BasicInfo)
	}()
}

func (cve *CveController) Upload(c *gin.Context) {
	var req uvpDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		commonctl.QueryFailure(c, err)

		return
	}

	cmd, err := req.toCmd()
	if err != nil {
		commonctl.Failure(c, err)

		return
	}

	if err = cve.repo.AddCVEBasicInfo(&cmd); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.SuccessCreate(c)
	}
}

func (cve *CveController) BasicInfo(c *gin.Context) {
	num, err := dp.NewCVENum(c.Param("cveNum"))
	if err != nil {
		commonctl.QueryFailure(c, err)

		return
	}

	if v, err := cve.repo.BasicInfo(num); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.SendRespGet(c, v)
	}
}
