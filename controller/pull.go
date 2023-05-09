package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/qinsheng99/go-domain-web/app"
	"github.com/qinsheng99/go-domain-web/common/api"
	commonctl "github.com/qinsheng99/go-domain-web/common/controller"
	"github.com/qinsheng99/go-domain-web/domain/elastic"
	_const "github.com/qinsheng99/go-domain-web/utils/const"
)

type BasePull struct {
	p app.PullServiceImpl
}

func AddRoutePull(r *gin.RouterGroup, pull elastic.RepoPullImpl) {
	basepull := BasePull{p: app.NewPullService(pull)}

	group := r.Group("/board")
	{
		group.GET("/refresh/:type", basepull.Refresh)
	}
}

func (b BasePull) Refresh(c *gin.Context) {
	var err error
	switch c.Param("type") {
	case _const.RefreshPr:
		err = b.p.Refresh(nil)
	case _const.RefreshIssue:
	case _const.RefreshRepo:
	default:
		commonctl.Failure(c, fmt.Errorf("unkonwn refresh type: %s", c.Param("type")))

		return
	}

	if err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.SuccessCreate(c)
	}
}

func (b BasePull) PRList(c *gin.Context) {
	var req api.RequestPull
	if err := c.ShouldBindWith(&req, binding.Form); err != nil {
		commonctl.Failure(c, err)

		return
	}

	req.SetDefault()

	if v, err := b.p.PullList(req, nil); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.Success(c, v)
	}
}
