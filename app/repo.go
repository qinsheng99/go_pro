package app

import (
	"github.com/qinsheng99/go-domain-web/domain"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/project/sort/domain/dp"
)

type repoService struct {
	p repository.RepoImpl
}

type RepoServiceImpl interface {
	RepoNames(p dp.Page, s dp.Size, name string) ([]domain.Repo, error)
	FindRepo(string) (*domain.Repo, error)
	FindRepoWith(id int) (domain.RepoWith, error)
}

func NewRepoService(p repository.RepoImpl) RepoServiceImpl {
	return &repoService{
		p: p,
	}
}

func (r repoService) RepoNames(p dp.Page, s dp.Size, name string) ([]domain.Repo, error) {
	return r.p.RepoNames(p, s, name)
}

func (r repoService) FindRepo(name string) (*domain.Repo, error) {
	return r.p.FindRepo(name)
}

func (r repoService) FindRepoWith(id int) (domain.RepoWith, error) {
	return r.p.FindRepoWith(id)
}
