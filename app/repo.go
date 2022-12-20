package app

import (
	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
)

type repoService struct {
	p repository.RepoImpl
}

type RepoServiceImpl interface {
	RepoNames(p api.Pages, name string) ([]mysql.Repo, error)
	FindRepo(string) (*mysql.Repo, error)
}

func NewRepoService(p repository.RepoImpl) RepoServiceImpl {
	return &repoService{
		p: p,
	}
}

func (r repoService) RepoNames(p api.Pages, name string) ([]mysql.Repo, error) {
	return r.p.RepoNames(p, name)
}

func (r repoService) FindRepo(name string) (*mysql.Repo, error) {
	return r.p.FindRepo(name)
}
