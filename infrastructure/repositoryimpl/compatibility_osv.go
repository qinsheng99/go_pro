package repositoryimpl

import (
	"encoding/json"
	"errors"

	"github.com/qinsheng99/go-domain-web/common/api"
	"github.com/qinsheng99/go-domain-web/domain"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/utils"
)

type repoOsv struct {
	url string
	req utils.ReqImpl
	cli dbImpl
}

func NewRepoOsv(url string, req utils.ReqImpl, cli dbImpl) repository.RepoOsvImpl {
	return &repoOsv{url: url, req: req, cli: cli}
}

func (r *repoOsv) SyncOsv() (string, error) {
	osvList, err := r.parserOsv()
	if err != nil {
		return "", err
	}

	if len(osvList) == 0 {
		return "", errors.New("resource data is nil")
	}

	if err = r.syncOsv(osvList); err != nil {
		return "", err
	}

	return "success", nil
}

func (r *repoOsv) syncOsv(osvList []api.Osv) error {
	var err error
	tx := r.cli.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
			return
		}

		tx.Rollback()
	}()

	for k := range osvList {
		v := osvList[k]
		filter := compatibilityOsvDO{OsVersion: v.OsVersion}
		if len(v.PlatformResult) == 0 && len(v.ToolsResult) == 0 {
			err = r.cli.DeleteTransaction(&filter, tx)
			if err != nil {
				return err
			}

			continue
		}
		var tools, platform []byte
		tools, err = json.Marshal(v.ToolsResult)
		if err != nil {
			return err
		}
		platform, err = json.Marshal(v.PlatformResult)
		if err != nil {
			return err
		}

		var do compatibilityOsvDO
		toCompatibilityOsvDO(&do, v, tools, platform)

		var ok bool
		if ok, err = r.cli.Exist(&filter, &compatibilityOsvDO{}); err == nil && ok {
			err = r.cli.UpdateTransaction(&filter, &do, tx)
			if err != nil {
				tx.Rollback()
				return err
			}
		} else if err == nil {
			if err = r.cli.InsertTransaction(&filter, &do, tx); err != nil {
				tx.Rollback()
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func (r *repoOsv) parserOsv() (osv []api.Osv, err error) {
	_, err = r.req.CustomRequest(
		r.url, "GET", nil, map[string]string{"Content-Type": "text/html"}, nil, true, &osv,
	)

	return
}

func (r *repoOsv) Find(opt domain.OsvOptions) (_ []domain.CompatibilityOsv, _ int64, _ error) {
	//r.cli.
	return nil, 0, nil
}
