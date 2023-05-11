package repositoryimpl

import (
	"encoding/json"
	"errors"

	"gorm.io/gorm"

	"github.com/qinsheng99/go-domain-web/common/api"
	"github.com/qinsheng99/go-domain-web/common/infrastructure/dao"
	"github.com/qinsheng99/go-domain-web/project/openbackend/domain"
	"github.com/qinsheng99/go-domain-web/project/openbackend/domain/repository"
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

	f := func(tx *gorm.DB) error {
		for k := range osvList {
			v := osvList[k]
			filter := compatibilityOsvDO{OsVersion: v.OsVersion}
			if len(v.PlatformResult) == 0 && len(v.ToolsResult) == 0 {
				err = r.cli.Delete(&filter, tx)
				if err != nil {
					return err
				}

				continue
			}
			var tools, platform []byte
			if tools, err = json.Marshal(v.ToolsResult); err != nil {
				return err
			}

			if platform, err = json.Marshal(v.PlatformResult); err != nil {
				return err
			}

			var do compatibilityOsvDO
			toCompatibilityOsvDO(&do, v, tools, platform)

			var ok bool
			var res compatibilityOsvDO
			if ok, err = r.cli.Exist(tx, &filter, &res); err != nil {
				return err
			} else if ok {
				do.Id = res.Id
			}

			if err = r.cli.CreateOrUpdate(tx, &do, osvUpdates...); err != nil {
				return err
			}
		}

		return nil
	}

	if err = r.cli.Transaction(nil, f); err != nil {
		return err
	}

	return nil
}

func (r *repoOsv) parserOsv() (osv []api.Osv, err error) {
	_, err = r.req.CustomRequest(
		r.url, "GET", nil, map[string]string{"Content-Type": "text/html"}, nil, true, &osv,
	)

	return
}

func (r *repoOsv) OsvList(opt domain.OsvOptions) ([]domain.CompatibilityOsvInfo, int64, error) {
	f := func(db *gorm.DB) *gorm.DB {
		if opt.KeyWord != "" {
			db.Where(
				db.
					Where("osv_name like ?", "%"+opt.KeyWord+"%").
					Or("os_version like ?", "%"+opt.KeyWord+"%").
					Or("type like ?", "%"+opt.KeyWord+"%"),
			)
		}
		if opt.OsvName != "" {
			db.Where("osv_name like ?", opt.OsvName)
		}

		if opt.Type != "" {
			db.Where("type = ?", opt.Type)
		}

		return db
	}

	total, err := r.cli.Count(nil, f)
	if err != nil {
		return nil, 0, err
	}

	var do []compatibilityOsvDO
	if err = r.cli.GetRecords(
		nil, f,
		&do,
		dao.Pagination{PageNum: opt.Page.Page(), CountPerPage: opt.Size.Size()},
		[]dao.SortByColumn{{Column: "id"}},
	); err != nil {
		return nil, 0, err
	}

	var res = make([]domain.CompatibilityOsvInfo, len(do))

	for i, v := range do {
		res[i] = v.toCompatibilityOsvInfo()
	}

	return res, total, nil
}
