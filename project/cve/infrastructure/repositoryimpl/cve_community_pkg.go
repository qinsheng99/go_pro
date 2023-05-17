package repositoryimpl

import (
	"gorm.io/gorm"

	"github.com/qinsheng99/go-domain-web/project/cve/domain"
)

const size = 200

type communityPkg struct {
	cli dbimpl
}

func (c communityPkg) AddApplicationPkg(app []domain.ApplicationPackage) error {
	res := c.toAppPkgDO(app)
	if len(res) == 0 {
		return nil
	}

	f := func(tx *gorm.DB) error {
		if err := c.cli.UpdateRecord(
			tx, &cveCommunityPkgDO{Community: res[0].Community}, &cveCommunityPkgDO{Status: pkgDelete},
		); err != nil && !c.cli.IsRowNotFound(err) {
			return err
		}
		for i := range res {
			v := &res[i]

			var do cveCommunityPkgDO
			err := c.cli.GetRecord(tx, func(db *gorm.DB) *gorm.DB {
				return db.Where(
					"package_name = ? and version = ? and community = ? and repo = ?",
					v.PackageName, v.Version, v.Community, v.Repo,
				)
			}, &do)
			if err == nil {
				v.Id = do.Id
				v.Status = pkgUpdate
			}
		}

		if err := c.cli.CreateOrUpdate(tx.Session(&gorm.Session{CreateBatchSize: 200}), res, pkgUpdates...); err != nil {
			return err
		}

		return c.cli.Delete(tx, &cveCommunityPkgDO{Status: "delete"})
	}

	return c.cli.Transaction(nil, f)
}

//
//func (c communityPkg) AddBasePkg(app []domain.BasePackage) error {
//	res := c.toAppPkgDO(app)
//
//	f := func(tx *gorm.DB) error {
//		if err := c.cli.UpdateRecord(
//			tx, &cveCommunityPkgDO{Community: app.Community.Community()}, &cveCommunityPkgDO{Status: "delete"},
//		); err != nil && !c.cli.IsRowNotFound(err) {
//			return err
//		}
//		for i := range res {
//			v := &res[i]
//
//			var do cveCommunityPkgDO
//			err := c.cli.GetRecord(tx, func(db *gorm.DB) *gorm.DB {
//				return db.Where(
//					"package_name = ? and version = ? and community = ? and repo = ?",
//					v.PackageName, v.Version, v.Community, v.Repo,
//				)
//			}, &do)
//			if err == nil {
//				v.Id = do.Id
//				v.Status = "update"
//			}
//		}
//
//		if err := c.cli.CreateOrUpdate(tx.Session(&gorm.Session{CreateBatchSize: 200}), res, pkgUpdates...); err != nil {
//			return err
//		}
//
//		return c.cli.Delete(tx, &cveCommunityPkgDO{Status: "delete"})
//	}
//
//	return c.cli.Transaction(nil, f)
//}
