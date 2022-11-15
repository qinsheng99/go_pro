package route

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-domain-web/app"
	"github.com/qinsheng99/go-domain-web/config"
	"github.com/qinsheng99/go-domain-web/controller"
	kubercontrol "github.com/qinsheng99/go-domain-web/controller/kubernetes"
	"github.com/qinsheng99/go-domain-web/docs"
	"github.com/qinsheng99/go-domain-web/infrastructure/kubernetes"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
	"github.com/qinsheng99/go-domain-web/infrastructure/redis"
	"github.com/qinsheng99/go-domain-web/infrastructure/repository"
	"github.com/qinsheng99/go-domain-web/infrastructure/sort"
	"github.com/qinsheng99/go-domain-web/utils"
	"github.com/qinsheng99/go-domain-web/utils/const"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func SetRoute(r *gin.Engine) {
	docs.SwaggerInfo.Description = "go-domain-web"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Host = "localhost:8000"
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := r.Group("/v1")

	controller.AddRouteOsv(
		group,
		repository.NewRepoOsv(_const.ParserOsvJsonFile, utils.NewRequest(nil),
			mysql.NewOsvMapper(),
		),
	)

	kubercontrol.AddRoutePod(group, kubernetes.NewPodImpl(config.Conf.KubernetesConfig))

	kubercontrol.AddRouteConfigMap(group, kubernetes.NewConfigImpl(config.Conf.KubernetesConfig))

	controller.AddRouteSort(group, app.NewSortService(sort.NewSort()))

	controller.AddRouteRedis(group, app.NewRedisService(redis.NewredisImpl(redis.GetRedis())))
}
