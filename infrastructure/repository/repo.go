package repository

import (
	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
)

type repoRepo struct {
	r mysql.RepoMapper
}

func NewRepoR(r mysql.RepoMapper) repository.RepoImpl {
	return repoRepo{r: r}
}

func (r repoRepo) RepoNames(p api.Pages) ([]string, error) {
	return r.r.RepoNames(p)
}

func (r repoRepo) FindRepo(name string) (*mysql.Repo, error) {
	return r.r.FindRepo(name)
}
