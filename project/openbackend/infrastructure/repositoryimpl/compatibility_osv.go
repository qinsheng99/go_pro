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

	total, err := r.cli.Count(f)
	if err != nil {
		return nil, 0, err
	}

	var do []compatibilityOsvDO
	if err = r.cli.GetRecords(
		f,
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
