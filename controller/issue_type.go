package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
	"github.com/qinsheng99/go-domain-web/utils"
)

type BaseIssueType struct {
}

func AddRouteIssueType(r *gin.RouterGroup) {
	baseIssue := &BaseIssueType{}

	group := r.Group("/issue-type")

	func() {
		group.GET("/list", baseIssue.List)
		group.GET("/one", baseIssue.One)
	}()
}

func (BaseIssueType) List(c *gin.Context) {
	var i mysql.IssueType
	list, err := i.List()
	if err != nil {
		utils.Failure(c, err)
		return
	}

	utils.Success(c, http.StatusOK, list)
}

func (BaseIssueType) One(c *gin.Context) {
	var i mysql.IssueType
	id, _ := strconv.Atoi(c.Query("unique"))
	i.UniqueId = int64(id)
	err := i.Find()
	if err != nil {
		utils.Failure(c, err)
		return
	}

	utils.Success(c, http.StatusOK, i)
}
