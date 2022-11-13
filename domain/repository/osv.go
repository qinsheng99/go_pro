package repository

import (
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
)

type RepoOsvImpl interface {
	SyncOsv() (string, error)
	Find() ([]mysql.OeCompatibilityOsv, int64, error)
}
