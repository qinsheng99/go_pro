package controller

import (
	"github.com/gin-gonic/gin"
)

type CveController struct {
}

func AddRouteCve(r *gin.RouterGroup) {
	ctl := &CveController{}

	group := r.Group("/cve")

	func() {
		group.POST("/upload", ctl.Upload)
	}()
}

func (c *CveController) Upload(ctx *gin.Context) {

}
