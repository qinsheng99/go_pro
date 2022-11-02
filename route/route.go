package route

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-domain-web/controller"
	"github.com/qinsheng99/go-domain-web/infrastructure/score"
	"net/http"
	"os"
)

func SetRoute(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})

	controller.AddRouteScore(
		r,
		score.NewScore(
			os.Getenv("EVALUATE"),
			os.Getenv("CALCULATE")),
	)
}
