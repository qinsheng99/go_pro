package repository

import (
	"github.com/qinsheng99/go-domain-web/domain"
)

type RepoOsvImpl interface {
	SyncOsv() (string, error)
	Find(domain.OsvOptions) ([]domain.CompatibilityOsv, int64, error)
}
