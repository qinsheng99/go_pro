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

func AddRoutePkg(r *gin.RouterGroup) {
	ctl := &PkgController{}

	group := r.Group("/pkg")

	func() {
		group.POST("/upload/:type", ctl.Upload)
	}()
}

func (p *PkgController) Upload(c *gin.Context) {
	switch c.Param("type") {
	case "application":
		go p.AddApplicationPkg(nil)
	default:
		commonctl.Failure(c, errors.New("invalid type"))
	}

	commonctl.SuccessCreate(c)
}
