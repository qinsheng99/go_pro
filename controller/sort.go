package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qinsheng99/go-domain-web/api"
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
	var req api.Sort
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		utils.QueryFailure(c)
		return
	}
	switch c.Param("type") {
	case "select":
		b.s.Select(req.Data)
	case "bubbling":
		b.s.Bubbling(req.Data)
	case "insert":
		b.s.Insert(req.Data)
	case "quick":
		b.s.Quick(req.Data)
	default:
		utils.Failure(c, fmt.Errorf("unknown typ %s", c.Param("type")))
		return
	}
	utils.Success(c, http.StatusOK, req.Data)
}
