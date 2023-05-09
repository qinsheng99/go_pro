package server

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
	kebectl "github.com/qinsheng99/go-domain-web/project/kubernetes/controller"
	openctl "github.com/qinsheng99/go-domain-web/project/openbackend/controller"
	"github.com/qinsheng99/go-domain-web/project/openbackend/infrastructure/repositoryimpl"
	sortapp "github.com/qinsheng99/go-domain-web/project/sort/app"
	sortctl "github.com/qinsheng99/go-domain-web/project/sort/controller"
	"github.com/qinsheng99/go-domain-web/project/sort/infrastructure/sort"
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

	openctl.AddRouteOsv(
		group,
		repositoryimpl.NewRepoOsv(_const.ParserOsvJsonFile, utils.NewRequest(nil),
			mysql.NewMqDao(cfg.Mysql.Table.CompatibilityOsv),
		),
	)

	pull := elastic.NewPullMapper(cfg.Es.Indexs.PullIndex)

	kebectl.AddRoutePod(group, kubernetes.NewPodImpl(cfg.Kubernetes))

	kebectl.AddRouteConfigMap(group, kubernetes.NewConfigImpl(cfg.Kubernetes))

	sortctl.AddRouteSort(group, sortapp.NewSortService(sort.NewSort()))

	controller.AddRouteRedis(group, app.NewRedisService(redis.NewredisImpl(redis.GetRedis())))

	controller.AddRoutePull(group, elasticsearch.NewRepoPull(pull, utils.NewRequest(nil)))

	//controller.AddRouteRepo(group, repositoryimpl.NewRepoR(mysql.NewRepoMapper()))
}
