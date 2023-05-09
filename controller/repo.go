package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/qinsheng99/go-domain-web/app"
	commonctl "github.com/qinsheng99/go-domain-web/common/controller"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/project/sort/domain/dp"
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

	if repo, err := b.r.RepoNames(dp.NewPage(page), dp.NewSize(size), c.DefaultQuery("name", "")); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.Success(c, repo)
	}
}

func (b BaseRepo) FindRepo(c *gin.Context) {
	name := c.Query("repo")

	if repo, err := b.r.FindRepo(name); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.Success(c, repo)
	}
}

func (b BaseRepo) FindRepoWith(c *gin.Context) {
	var id = struct {
		Id int `json:"id" form:"id"`
	}{}

	if err := c.ShouldBindQuery(&id); err != nil {
		commonctl.Failure(c, err)

		return
	}

	if repo, err := b.r.FindRepoWith(id.Id); err != nil {
		commonctl.Failure(c, err)
	} else {
		commonctl.Success(c, repo)
	}
}
