package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/app"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/utils"
	"net/http"
	"strconv"
)

type BaseRepo struct {
	r app.RepoServiceImpl
}

func AddRouteRepo(r *gin.RouterGroup, repo repository.RepoImpl) {
	baserepo := &BaseRepo{r: app.NewRepoService(repo)}

	group := r.Group("/repo")

	func() {
		group.GET("/repo-names", baserepo.RepoNames)
		group.GET("/repo", baserepo.FindRepo)
	}()
}

func (b BaseRepo) RepoNames(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	var p = api.Pages{
		Page: page,
		Size: size,
	}

	repo, err := b.r.RepoNames(p)
	if err != nil {
		utils.Failure(c, err)
		return
	}

	utils.Success(c, http.StatusOK, repo)
}

func (b BaseRepo) FindRepo(c *gin.Context) {
	name := c.Query("repo")

	repo, err := b.r.FindRepo(name)
	if err != nil {
		utils.Failure(c, err)
		return
	}

	utils.Success(c, http.StatusOK, repo)
}
