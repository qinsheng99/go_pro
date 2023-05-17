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
	switch c.Param("type") {
	case "application":
		p.uploadApp(c)
	case "base":
		p.uploadBase(c)
	default:
		commonctl.Failure(c, errors.New("invalid type"))
	}
}

func (p *PkgController) uploadApp(c *gin.Context) {
	var req applicationPkgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		commonctl.QueryFailure(c, err)

		return
	}

	cmd, err := req.toApplicationPkgCmd()
	if err != nil {
		commonctl.Failure(c, err)

		return
	}

	if err = p.AddApplicationPkg(cmd); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.SuccessCreate(c)
	}
}

func (p *PkgController) uploadBase(c *gin.Context) {
	var req pkgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		commonctl.QueryFailure(c, err)

		return
	}

	cmd, err := req.toBasePkgCmd()
	if err != nil {
		commonctl.Failure(c, err)

		return
	}

	if err = p.AddBasePkg(cmd); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.SuccessCreate(c)
	}

}
