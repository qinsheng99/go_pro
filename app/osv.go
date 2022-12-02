package app

import (
	"encoding/json"

	"github.com/qinsheng99/go-domain-web/api/osv"
	"github.com/qinsheng99/go-domain-web/domain"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	_const "github.com/qinsheng99/go-domain-web/utils/const"
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
	Find() (*domain.ResultOsv, error)
}

func (o *osvService) SyncOsv() (string, error) {
	return o.osv.SyncOsv()
}

func (o *osvService) Find() (_ *domain.ResultOsv, _ error) {
	list, total, err := o.osv.Find()
	if err != nil {
		return nil, err
	}
	data := make([]domain.ROeCompatibilityOsv, 0, len(list))
	for _, v := range list {
		var t []osv.Record
		_ = json.Unmarshal([]byte(v.ToolsResult), &t)
		var p []osv.Record
		_ = json.Unmarshal([]byte(v.PlatformResult), &p)
		data = append(data, domain.ROeCompatibilityOsv{
			OeCompatibilityOsv: v,
			ToolsResult:        t,
			PlatformResult:     p,
			Updateime:          v.Updateime.Format(_const.Format),
		})
	}
	return &domain.ResultOsv{OsvList: data, Total: int64(total)}, nil
}
