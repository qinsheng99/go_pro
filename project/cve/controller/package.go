package controller

import (
	"errors"

	"github.com/gin-gonic/gin"

	commonctl "github.com/qinsheng99/go-domain-web/common/controller"
	"github.com/qinsheng99/go-domain-web/project/cve/app"
)

type PkgController struct {
	app.PkgService
}

func AddRoutePkg(r *gin.RouterGroup, service app.PkgService) {
	ctl := &PkgController{service}

	group := r.Group("/pkg")

	func() {
		group.POST("/upload/:type", ctl.Upload)
	}()
}

func (p *PkgController) Upload(c *gin.Context) {
	var req PkgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		commonctl.QueryFailure(c, err)

		return
	}

	switch c.Param("type") {
	case "application":
		cmd, err := req.toApplicationPkgCmd()
		if err != nil {
			commonctl.Failure(c, err)

			return
		}
		go p.AddApplicationPkg(&cmd)
	default:
		commonctl.Failure(c, errors.New("invalid type"))

		return
	}

	commonctl.SuccessCreate(c)
}
