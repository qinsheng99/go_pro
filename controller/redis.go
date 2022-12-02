package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/qinsheng99/go-domain-web/app"
	"github.com/qinsheng99/go-domain-web/utils"
)

type BaseRedis struct {
	r app.RedisServiceImpl
}

func AddRouteRedis(r *gin.RouterGroup, impl app.RedisServiceImpl) {
	baseSort := BaseRedis{r: impl}

	group := r.Group("/redis")

	func() {
		group.GET("/zadd", baseSort.Zadd)
		group.GET("/zrange", baseSort.Zrange)
		group.DELETE("/delete/:key", baseSort.Delete)
	}()
}

func (b *BaseRedis) Zadd(c *gin.Context) {
	var a = []int{123, 456, 789, 100, 23}
	var data []*redis.Z
	var score = 5000

	for k, v := range a {
		data = append(data, &redis.Z{
			Score:  float64(score - k),
			Member: "店铺号:" + strconv.Itoa(v),
		})
	}
	res, err := b.r.Zadd(context.Background(), "score", data...)
	if err != nil {
		utils.Failure(c, err)
		return
	}
	utils.Success(c, http.StatusOK, res)
}

func (b *BaseRedis) Zrange(c *gin.Context) {
	revrange, err := b.r.ZRevrange(context.Background(), "score", 0, -1)
	if err != nil {
		utils.Failure(c, err)
	}
	utils.Success(c, http.StatusOK, revrange)
}

// Delete
// @Title Delete
// @Description redis del
// @Tags redis
// @Accept application/json
// @Param key query string true "redis key"
// @Success 200 {integer} integer
// @Failure 500 system_error system error
// @Router /del/:key [delete]
func (b *BaseRedis) Delete(c *gin.Context) {
	del, err := b.r.Del(context.Background(), c.Param("key"))
	if err != nil {
		utils.Failure(c, err)
	}
	utils.Success(c, http.StatusOK, del)
}
