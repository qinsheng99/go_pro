package repository

import (
	"github.com/qinsheng99/go-domain-web/domain/dp"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
)

type RepoImpl interface {
	RepoNames(p dp.Page, s dp.Size, name string) ([]mysql.Repo, error)
	FindRepo(string) (*mysql.Repo, error)
	FindRepoWith(id int) (mysql.RepoWith, error)
}
