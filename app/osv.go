package app

import (
	"encoding/json"
	"github.com/qinsheng99/go-domain-web/api/osv_api"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/utils"
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
	Find() (*repository.ResultOsv, error)
}

func (o *osvService) SyncOsv() (string, error) {
	return o.osv.SyncOsv()
}

func (o *osvService) Find() (_ *repository.ResultOsv, _ error) {
	list, total, err := o.osv.Find()
	if err != nil {
		return nil, err
	}
	data := make([]repository.ROeCompatibilityOsv, 0, len(list))
	for _, v := range list {
		var t []osv_api.Record
		_ = json.Unmarshal([]byte(v.ToolsResult), &t)
		var p []osv_api.Record
		_ = json.Unmarshal([]byte(v.PlatformResult), &p)
		data = append(data, repository.ROeCompatibilityOsv{
			OeCompatibilityOsv: v,
			ToolsResult:        t,
			PlatformResult:     p,
			Updateime:          v.Updateime.Format(utils.Format),
		})
	}
	return &repository.ResultOsv{OsvList: data, Total: int64(total)}, nil
}
