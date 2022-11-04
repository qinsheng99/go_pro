package repository

import (
	"errors"
	"github.com/qinsheng99/go-domain-web/api/osv_api"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
	"github.com/qinsheng99/go-domain-web/utils"
)

type repoOsv struct {
	osv mysql.OsvMapper
	url string
	req utils.ReqImpl
}

func NewRepoOsv(url string, req utils.ReqImpl, osv mysql.OsvMapper) repository.RepoOsvImpl {
	return &repoOsv{url: url, req: req, osv: osv}
}

func (r *repoOsv) SyncOsv() (string, error) {
	osvList, err := r.parserOsv()
	if err != nil {
		return "", err
	}

	if len(osvList) == 0 {
		return "", errors.New("resource data is nil")
	}

	err = r.osv.SyncOsv(osvList)
	if err != nil {
		return "", err
	}

	return "success", nil
}

func (r *repoOsv) Find() (data []repository.ROeCompatibilityOsv, _ int64, _ error) {
	list, total, err := r.osv.OSVFindAll(osv_api.RequestOsv{})
	if err != nil {
		return nil, 0, err
	}

	data = make([]repository.ROeCompatibilityOsv, len(list))

	for _, v := range list {
		data = append(data, repository.ROeCompatibilityOsv{OeCompatibilityOsv: v})
	}
	return data, total, nil
}

func (r *repoOsv) parserOsv() (osv []osv_api.Osv, err error) {
	_, err = r.req.CustomRequest(
		r.url,
		"GET",
		nil,
		map[string]string{"Content-Type": "text/html"},
		nil,
		true,
		&osv)
	if err != nil {
		return nil, err
	}
	return
}
