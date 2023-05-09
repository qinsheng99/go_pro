package repository

import (
	"github.com/qinsheng99/go-domain-web/project/openbackend/domain"
)

type RepoOsvImpl interface {
	SyncOsv() (string, error)
	OsvList(domain.OsvOptions) ([]domain.CompatibilityOsvInfo, int64, error)
}
