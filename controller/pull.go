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
	_const "github.com/qinsheng99/go-domain-web/utils/const"
	"k8s.io/apimachinery/pkg/util/sets"
)

type BasePull struct {
	p app.PullServiceImpl
	base
}

func AddRoutePull(r *gin.RouterGroup, pull repository.RepoPullImpl) {
	basepull := BasePull{p: app.NewPullService(pull)}

	group := r.Group("/board")
	{
		group.GET("/refresh/:type", basepull.Refresh)
		group.GET("/pulls", basepull.PRList)
		//group.GET("/pulls/:field", basepull.PullField)
		group.GET("/pulls/authors", basepull.PullAuthors)
	}
}

var fields = sets.NewString(_const.PullsAuthors, _const.PullsAssignees, _const.PullsLabels, _const.PullsRef, _const.PullsSig, _const.PullsRepos)

func (b BasePull) Refresh(c *gin.Context) {
	var err error
	switch c.Param("type") {
	case _const.RefreshPr:
		err = b.p.Refresh(nil)
	case _const.RefreshIssue:

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

func (b BasePull) PullField(c *gin.Context) {
	var req api.RequestPull
	var err error
	if err = c.ShouldBindWith(&req, binding.Form); err != nil {
		utils.Failure(c, err)
		return
	}

	req.SetDefault()
	var res []string
	var total int64
	field := c.Param("field")
	switch fields.Has(field) {
	case true:
		total, res, err = b.p.PullFields(req, field)
	default:
		utils.Failure(c, fmt.Errorf("unkonwn pulls field: %s", field))
		return
	}

	if err != nil {
		utils.Failure(c, err)
		return
	}

	c.JSON(http.StatusOK, b.base.Response(res, req.Page, req.PerPage, int(total)))
}

func (b BasePull) PullAuthors(c *gin.Context) {
	var req api.RequestPull
	var err error
	if err = c.ShouldBindWith(&req, binding.Form); err != nil {
		utils.Failure(c, err)
		return
	}

	req.SetDefault()
	var res []string
	var total int64
	total, res, err = b.p.PullAuthors(req)

	if err != nil {
		utils.Failure(c, err)
		return
	}

	c.JSON(http.StatusOK, b.base.Response(res, req.Page, req.PerPage, int(total)))
}
