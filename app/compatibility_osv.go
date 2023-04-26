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
	List(domain.OsvOptions) (*CompatibilityOsvDTO, error)
}

func (o *osvService) SyncOsv() (string, error) {
	return o.osv.SyncOsv()
}

func (o *osvService) List(opt domain.OsvOptions) (*CompatibilityOsvDTO, error) {
	list, total, err := o.osv.OsvList(opt)
	if err != nil {
		return nil, err
	}

	return toCompatibilityOsvDTO(list, total), nil
}
