package app

import (
	"encoding/json"

	"github.com/qinsheng99/go-domain-web/api"
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
	Find(domain.OsvDP) (*ResultOsv, error)
}

func (o *osvService) SyncOsv() (string, error) {
	return o.osv.SyncOsv()
}

func (o *osvService) Find(osv domain.OsvDP) (_ *ResultOsv, _ error) {
	list, total, err := o.osv.Find(osv)
	if err != nil {
		return nil, err
	}
	data := make([]ROeCompatibilityOsv, 0, len(list))
	for _, v := range list {
		var t []api.Record
		_ = json.Unmarshal([]byte(v.ToolsResult), &t)
		var p []api.Record
		_ = json.Unmarshal([]byte(v.PlatformResult), &p)
		data = append(data, ROeCompatibilityOsv{
			OeCompatibilityOsv: v,
			ToolsResult:        t,
			PlatformResult:     p,
			Updateime:          v.Updateime.Format(_const.Format),
		})
	}
	return &ResultOsv{OsvList: data, Total: int64(total)}, nil
}
