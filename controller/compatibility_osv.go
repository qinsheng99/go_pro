package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/qinsheng99/go-domain-web/app"
	commonctl "github.com/qinsheng99/go-domain-web/common/controller"
	"github.com/qinsheng99/go-domain-web/common/logger"
	"github.com/qinsheng99/go-domain-web/domain/repository"
)

type BaseOsv struct {
	osv app.OsvServiceImpl
}

func AddRouteOsv(r *gin.RouterGroup, osv repository.RepoOsvImpl) {
	baseosv := &BaseOsv{osv: app.NewOsvService(osv)}

	group := r.Group("/osv")

	func() {
		group.GET("/syncOsv", baseosv.SyncOsv)
		group.POST("/list", baseosv.List)
	}()
}

func (b *BaseOsv) SyncOsv(c *gin.Context) {
	if result, err := b.osv.SyncOsv(); err != nil {
		logger.Log.Error("syncOsv failed :", err)
		commonctl.Failure(c, fmt.Errorf("syncOsv failed. An exception occurred."+result+err.Error()))
	} else {
		commonctl.Success(c, result)
	}

}

// List
// @Summary osv list
// @Description osv list
// @Tags  OSV
// @Accept json
// @Param	param  body	 requestOsv	 true	"body of get osv list"
// @Success 200 {object} app.CompatibilityOsvDTO
// @Failure 400 {object} ResponseData
// @Router /v1/osv/list [post]
func (b *BaseOsv) List(c *gin.Context) {
	var osv requestOsv
	if err := c.ShouldBindJSON(&osv); err != nil {
		commonctl.QueryFailure(c, err)

		return
	}

	o := osv.tocmd()

	if result, err := b.osv.List(o); err != nil {
		logger.Log.Error("syncOsv failed :", err)
		commonctl.Failure(c, err)
	} else {
		commonctl.Success(c, result)
	}
}
