package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/app"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/utils"
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
		group.GET("/repo-with", baserepo.FindRepoWith)
	}()
}

func (b BaseRepo) RepoNames(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	var p = api.Pages{
		Page: page,
		Size: size,
	}

	repo, err := b.r.RepoNames(p, c.DefaultQuery("name", ""))
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

func (b BaseRepo) FindRepoWith(c *gin.Context) {
	var id = struct {
		Id int `json:"id" form:"id"`
	}{}

	if err := c.ShouldBindQuery(&id); err != nil {
		utils.Failure(c, err)
		return
	}

	repo, err := b.r.FindRepoWith(id.Id)
	if err != nil {
		utils.Failure(c, err)
		return
	}

	utils.Success(c, http.StatusOK, repo)
}
