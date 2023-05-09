package repository

import (
	"github.com/qinsheng99/go-domain-web/domain"
	"github.com/qinsheng99/go-domain-web/project/sort/domain/dp"
)

type RepoImpl interface {
	RepoNames(p dp.Page, s dp.Size, name string) ([]domain.Repo, error)
	FindRepo(string) (*domain.Repo, error)
	FindRepoWith(id int) (domain.RepoWith, error)
}
