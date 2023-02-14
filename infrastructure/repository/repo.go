package repository

import (
	"github.com/qinsheng99/go-domain-web/domain/dp"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
)

type repoRepo struct {
	r mysql.RepoMapper
}

func NewRepoR(r mysql.RepoMapper) repository.RepoImpl {
	return repoRepo{r: r}
}

func (r repoRepo) RepoNames(p dp.Page, s dp.Size, name string) ([]mysql.Repo, error) {
	return r.r.RepoNames(p, s, name)
}

func (r repoRepo) FindRepo(name string) (*mysql.Repo, error) {
	return r.r.FindRepo(name)
}

func (r repoRepo) FindRepoWith(id int) (mysql.RepoWith, error) {
	return r.r.FindRepoWith(id)
}
