package repository

import (
	"github.com/qinsheng99/go-domain-web/domain"
)

type RepoOsvImpl interface {
	SyncOsv() (string, error)
	Find(domain.OsvDP) ([]domain.OeCompatibilityOsv, int64, error)
}
