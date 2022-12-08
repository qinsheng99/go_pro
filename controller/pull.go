package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/app"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/utils"
)

type BasePull struct {
	p app.PullServiceImpl
}

func AddRoutePull(r *gin.RouterGroup, pull repository.RepoPullImpl) {
	basepull := BasePull{p: app.NewPullService(pull)}

	group := r.Group("/board")
	{
		group.GET("/refresh/:type", basepull.Refresh)
		group.GET("/pulls", basepull.PRList)
		group.GET("/pulls/:type", basepull.PullField)
	}
}

func (b BasePull) Refresh(c *gin.Context) {
	var err error
	switch c.Param("type") {
	case "pr":
		err = b.p.Refresh(nil)
	case "issue":

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

	utils.Success(c, http.StatusOK, api.ResponsePull{Total: total, Data: list})
}

func (b BasePull) PullField(c *gin.Context) {
	var req api.RequestPull
	var err error
	if err = c.ShouldBindWith(&req, binding.Form); err != nil {
		utils.Failure(c, err)
		return
	}

	req.SetDefault()
	var res []string
	switch c.Param("type") {
	case "author":
		res, err = b.p.PullAuthor(req, nil)
	default:
		utils.Failure(c, fmt.Errorf("unkonwn pulls type: %s", c.Param("type")))
		return
	}

	if err != nil {
		utils.Failure(c, err)
		return
	}

	utils.Success(c, http.StatusOK, res)
}
