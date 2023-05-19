package controller

import (
	"errors"

	"github.com/gin-gonic/gin"

	commonctl "github.com/qinsheng99/go-domain-web/common/controller"
	"github.com/qinsheng99/go-domain-web/project/cve/app"
)

const (
	application = "application"
	base        = "base"
)

var invalidType = errors.New("invalid type")

type PkgController struct {
	app.PkgService
}

func AddRoutePkg(r *gin.RouterGroup, service app.PkgService) {
	ctl := &PkgController{service}

	group := r.Group("/pkg")

	func() {
		group.POST("/:type/upload", ctl.Upload)

		group.GET("/:type/list", ctl.List)
	}()
}

func (p *PkgController) Upload(c *gin.Context) {
	switch c.Param("type") {
	case application:
		p.uploadApp(c)
	case base:
		p.uploadBase(c)
	default:
		commonctl.Failure(c, invalidType)
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

	if err = p.AddApplicationPkg(&cmd); err != nil {
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

	if err = p.AddBasePkg(&cmd); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.SuccessCreate(c)
	}
}

func (p *PkgController) List(c *gin.Context) {
	var q pkgListQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		commonctl.QueryFailure(c, err)

		return
	}

	switch c.Param("type") {
	case base:
		p.listBase(q, c)
	case application:
		p.listApplication(q, c)
	default:
		commonctl.Failure(c, invalidType)
	}

}

func (p *PkgController) listBase(q pkgListQuery, c *gin.Context) {
	opt, err := q.toOptFindPkgs()
	if err != nil {
		commonctl.Failure(c, err)

		return
	}

	if pkgs, err := p.PkgService.ListBasePkgs(opt); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.SendRespGet(c, pkgs)
	}
}

func (p *PkgController) listApplication(q pkgListQuery, c *gin.Context) {
	opt, err := q.toOptFindPkgs()
	if err != nil {
		commonctl.Failure(c, err)

		return
	}

	if pkgs, err := p.PkgService.ListApplicationPkgs(opt); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.SendRespGet(c, pkgs)
	}
}
