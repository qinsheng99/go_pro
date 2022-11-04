package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-domain-web/app"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/logger"
	"github.com/qinsheng99/go-domain-web/utils"
)

type BaseOsv struct {
	osv app.OsvServiceImpl
}

func AddRouteOsv(r *gin.RouterGroup, osv repository.RepoOsvImpl) {
	baseosv := BaseOsv{osv: app.NewOsvService(osv)}

	group := r.Group("/osv")

	func() {
		group.GET("/syncOsv", baseosv.SyncOsv)
		group.POST("/find", baseosv.Find)
	}()
}

func (b *BaseOsv) SyncOsv(c *gin.Context) {
	result, err := b.osv.SyncOsv()
	if err != nil {
		logger.Log.Error("syncOsv failed :", err)
		utils.Failure(c, fmt.Errorf("syncOsv failed. An exception occurred."+result+err.Error()))
		return
	}
	utils.Success(c, result)
}

func (b *BaseOsv) Find(c *gin.Context) {
	result, _, err := b.osv.Find()
	if err != nil {
		logger.Log.Error("syncOsv failed :", err)
		utils.Failure(c, err)
		return
	}
	utils.Success(c, result)
}