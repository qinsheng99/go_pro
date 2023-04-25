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
	Find(domain.OsvDP) (*resultOsvDTO, error)
}

func (o *osvService) SyncOsv() (string, error) {
	return o.osv.SyncOsv()
}

func (o *osvService) Find(osv domain.OsvDP) (*resultOsvDTO, error) {
	list, total, err := o.osv.Find(osv)
	if err != nil {
		return nil, err
	}

	return toResultOsvDTO(list, total), nil
}
