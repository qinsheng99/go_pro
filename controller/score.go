package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qinsheng99/go-domain-web/api/score_api"
	"github.com/qinsheng99/go-domain-web/app"
	"github.com/qinsheng99/go-domain-web/domain/score"
	"github.com/qinsheng99/go-py/sdk"
	"net/http"
)

func AddRouteScore(r *gin.RouterGroup, s score.Score) {
	baseScore := &BaseScore{
		s:   app.NewScoreService(s),
		cli: sdk.NewCalculateEvaluate("https://xihescript.test.osinfra.cn"),
	}

	group := r.Group("/score")

	func() {
		group.POST("/evaluate", baseScore.Evaluate)
		group.POST("/calculate", baseScore.Calculate)
		group.POST("/sdk-calculate", baseScore.SdkCalculate)
		group.POST("/sdk-evaluate", baseScore.SdkEvaluate)
	}()

}

type BaseScore struct {
	s   app.ScoreServiceImpl
	cli sdk.CalculateEvaluate
}

func (b *BaseScore) Evaluate(c *gin.Context) {
	col := score_api.Score{}
	if err := c.ShouldBindBodyWith(&col, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var res score_api.ScoreRes
	err := b.s.Evaluate(col, &res)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (b *BaseScore) Calculate(c *gin.Context) {
	col := score_api.Score{}
	if err := c.ShouldBindBodyWith(&col, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var res score_api.ScoreRes
	err := b.s.Calculate(col, &res)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (b *BaseScore) SdkEvaluate(c *gin.Context) {
	col := sdk.Calculate{}
	if err := c.ShouldBindBodyWith(&col, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	b.cli.Evaluate(&col)
}

func (b *BaseScore) SdkCalculate(c *gin.Context) {
	col := sdk.Calculate{}
	if err := c.ShouldBindBodyWith(&col, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	b.cli.Calculate(&col)
}
