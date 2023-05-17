package repositoryimpl

import (
	"gorm.io/gorm"

	"github.com/qinsheng99/go-domain-web/project/cve/domain"
)

const size = 200

type communityPkg struct {
	db dbimpl
}

func (c communityPkg) AddApplicationPkg(app []domain.ApplicationPackage) error {
	res := c.toAppPkgDO(app)
	if len(res) == 0 {
		return nil
	}

	return c.db.Transaction(c.db.DB(), c.transactionF(res))
}

func (c communityPkg) AddBasePkg(app []domain.BasePackage) error {
	res := c.toBasePkgDO(app)
	if len(res) == 0 {
		return nil
	}

	return c.db.Transaction(c.db.DB(), c.transactionF(res))
}

func (c communityPkg) transactionF(res []communityPkgDO) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		if err := c.db.UpdateRecord(
			tx, &communityPkgDO{Community: res[0].Community}, &communityPkgDO{Status: pkgDelete},
		); err != nil && !c.db.IsRowNotFound(err) {
			return err
		}
		for i := range res {
			v := &res[i]
			var do communityPkgDO
			if err := c.db.GetRecord(
				tx,
				func(db *gorm.DB) *gorm.DB {
					return db.Where("package_name = ? and version = ? and community = ? and repo = ?",
						v.PackageName, v.Version, v.Community, v.Repo,
					)
				},
				&do,
			); err == nil {
				v.Id = do.Id
				v.Status = pkgUpdate
			}
		}

		if err := c.db.CreateOrUpdate(tx.Session(&gorm.Session{CreateBatchSize: size}), res, pkgUpdates...); err != nil {
			return err
		}

		return c.db.Delete(tx, &communityPkgDO{Status: pkgDelete})
	}
}
