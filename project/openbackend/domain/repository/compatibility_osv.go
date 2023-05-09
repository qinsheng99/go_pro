package repository

import (
	"github.com/qinsheng99/go-domain-web/domain"
)

type RepoOsvImpl interface {
	SyncOsv() (string, error)
	OsvList(domain.OsvOptions) ([]domain.CompatibilityOsvInfo, int64, error)
}
