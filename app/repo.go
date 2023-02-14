package app

import (
	"github.com/qinsheng99/go-domain-web/domain/dp"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
)

type repoService struct {
	p repository.RepoImpl
}

type RepoServiceImpl interface {
	RepoNames(p dp.Page, s dp.Size, name string) ([]mysql.Repo, error)
	FindRepo(string) (*mysql.Repo, error)
	FindRepoWith(id int) (mysql.RepoWith, error)
}

func NewRepoService(p repository.RepoImpl) RepoServiceImpl {
	return &repoService{
		p: p,
	}
}

func (r repoService) RepoNames(p dp.Page, s dp.Size, name string) ([]mysql.Repo, error) {
	return r.p.RepoNames(p, s, name)
}

func (r repoService) FindRepo(name string) (*mysql.Repo, error) {
	return r.p.FindRepo(name)
}

func (r repoService) FindRepoWith(id int) (mysql.RepoWith, error) {
	return r.p.FindRepoWith(id)
}
