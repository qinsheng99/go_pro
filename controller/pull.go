package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/app"
	"github.com/qinsheng99/go-domain-web/domain/elastic"
	"github.com/qinsheng99/go-domain-web/utils"
	_const "github.com/qinsheng99/go-domain-web/utils/const"
)

type BasePull struct {
	p app.PullServiceImpl
	base
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
		utils.Failure(c, fmt.Errorf("unkonwn refresh type: %s", c.Param("type")))
		return
	}

	if err != nil {
		utils.Failure(c, err)
		return
	}
	utils.Success(c, http.StatusCreated, "")
}

func (b BasePull) PRList(c *gin.Context) {
	var req api.RequestPull
	if err := c.ShouldBindWith(&req, binding.Form); err != nil {
		utils.Failure(c, err)
		return
	}

	req.SetDefault()

	list, total, err := b.p.PullList(req, nil)
	if err != nil {
		utils.Failure(c, err)
		return
	}

	c.JSON(http.StatusOK, b.base.Response(list, req.Page, req.PerPage, int(total)))
}
