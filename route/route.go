package route

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-domain-web/config"
	"github.com/qinsheng99/go-domain-web/controller"
	kubercontrol "github.com/qinsheng99/go-domain-web/controller/kubernetes"
	"github.com/qinsheng99/go-domain-web/infrastructure/kubernetes"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
	"github.com/qinsheng99/go-domain-web/infrastructure/repository"
	"github.com/qinsheng99/go-domain-web/infrastructure/score"
	"github.com/qinsheng99/go-domain-web/utils"
	"net/http"
	"os"
)

func SetRoute(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})

	group := r.Group("/v1")

	controller.AddRouteScore(
		group,
		score.NewScore(
			os.Getenv("EVALUATE"),
			os.Getenv("CALCULATE")),
	)

	controller.AddRouteOsv(
		group,
		repository.NewRepoOsv(utils.ParserOsvJsonFile, utils.NewRequest(nil),
			mysql.NewOsvMapper(),
		),
	)

	kubercontrol.AddRoutePod(group, kubernetes.NewPodImpl(config.Conf.PodConfig))
}
