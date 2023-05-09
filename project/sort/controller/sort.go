package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	commonctl "github.com/qinsheng99/go-domain-web/common/controller"
	"github.com/qinsheng99/go-domain-web/project/sort/app"
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
		commonctl.QueryFailure(c, err)

		return
	}

	s, err := req.tocmd()
	if err != nil {
		commonctl.Failure(c, err)

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
		commonctl.Failure(c, fmt.Errorf("unknown typ %s", c.Param("type")))

		return
	}

	commonctl.Success(c, s.Fields.SortField())
}
