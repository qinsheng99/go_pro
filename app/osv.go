package app

import (
	"github.com/qinsheng99/go-domain-web/domain/repository"
)

type osvService struct {
	osv repository.RepoOsvImpl
}

func NewOsvService(osv repository.RepoOsvImpl) OsvServiceImpl {
	return &osvService{
		osv: osv,
	}
}

type OsvServiceImpl interface {
	SyncOsv() (string, error)
	Find() ([]repository.ROeCompatibilityOsv, int64, error)
}

func (o *osvService) SyncOsv() (string, error) {
	return o.osv.SyncOsv()
}

func (o *osvService) Find() ([]repository.ROeCompatibilityOsv, int64, error) {
	return o.osv.Find()
}
