package repository

import (
	"github.com/qinsheng99/go-domain-web/domain"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
)

type RepoOsvImpl interface {
	SyncOsv() (string, error)
	Find(domain.OsvDP) ([]mysql.OeCompatibilityOsv, int64, error)
}
