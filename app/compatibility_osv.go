package app

import (
	"github.com/qinsheng99/go-domain-web/domain"
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
	Find(domain.OsvOptions) (*compatibilityOsvDTO, error)
}

func (o *osvService) SyncOsv() (string, error) {
	return o.osv.SyncOsv()
}

func (o *osvService) Find(opt domain.OsvOptions) (*compatibilityOsvDTO, error) {
	list, total, err := o.osv.Find(opt)
	if err != nil {
		return nil, err
	}

	return toCompatibilityOsvDTO(list, total), nil
}
