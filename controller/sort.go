package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/qinsheng99/go-domain-web/app"
	"github.com/qinsheng99/go-domain-web/utils"
)

type BaseSort struct {
	s app.SortServiceImpl
}

func AddRouteSort(r *gin.RouterGroup, s app.SortServiceImpl) {
	baseSort := BaseSort{s: s}

	group := r.Group("/sort")

	func() {
		group.POST("/:type", baseSort.SelectSort)
	}()
}

func (b *BaseSort) SelectSort(c *gin.Context) {
	var req Sort
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		utils.QueryFailure(c)
		return
	}

	s, err := req.tocmd()
	if err != nil {
		utils.Failure(c, err)
		return
	}

	switch c.Param("type") {
	case "select":
		b.s.Select(s)
	case "bubbling":
		b.s.Bubbling(s)
	case "insert":
		b.s.Insert(s)
	case "quick":
		b.s.Quick(s)
	default:
		utils.Failure(c, fmt.Errorf("unknown typ %s", c.Param("type")))
		return
	}

	utils.Success(c, http.StatusOK, s.Fields.SortField())
}
