package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/qinsheng99/go-domain-web/app"
	"github.com/qinsheng99/go-domain-web/common/infrastructure/elastic"
	"github.com/qinsheng99/go-domain-web/common/infrastructure/mysql"
	"github.com/qinsheng99/go-domain-web/config"
	"github.com/qinsheng99/go-domain-web/controller"
	"github.com/qinsheng99/go-domain-web/docs"
	"github.com/qinsheng99/go-domain-web/infrastructure/elasticsearch"
	"github.com/qinsheng99/go-domain-web/infrastructure/kubernetes"
	"github.com/qinsheng99/go-domain-web/infrastructure/redis"
	"github.com/qinsheng99/go-domain-web/infrastructure/repositoryimpl"
	"github.com/qinsheng99/go-domain-web/infrastructure/sort"
	"github.com/qinsheng99/go-domain-web/utils"
	_const "github.com/qinsheng99/go-domain-web/utils/const"
)

func SetRoute(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Description = "go-domain-web"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Host = "8000"
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := r.Group("/v1")

	controller.AddRouteOsv(
		group,
		repositoryimpl.NewRepoOsv(
			_const.ParserOsvJsonFile,
			utils.NewRequest(nil),
			mysql.NewMqDao(cfg.MysqlConfig.Table.CompatibilityOsv),
		),
	)

	pull := elastic.NewPullMapper(cfg.EsConfig.Indexs.PullIndex)

	controller.AddRoutePod(group, kubernetes.NewPodImpl(cfg.KubernetesConfig))

	controller.AddRouteConfigMap(group, kubernetes.NewConfigImpl(cfg.KubernetesConfig))

	controller.AddRouteSort(group, app.NewSortService(sort.NewSort()))

	controller.AddRouteRedis(group, app.NewRedisService(redis.NewredisImpl(redis.GetRedis())))

	controller.AddRoutePull(group, elasticsearch.NewRepoPull(pull, utils.NewRequest(nil)))

	//controller.AddRouteRepo(group, repositoryimpl.NewRepoR(mysql.NewRepoMapper()))
}
