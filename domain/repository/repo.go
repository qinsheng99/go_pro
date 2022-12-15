package repository

import (
	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
)

type RepoImpl interface {
	RepoNames(p api.Pages) ([]string, error)
	FindRepo(string) (*mysql.Repo, error)
}
