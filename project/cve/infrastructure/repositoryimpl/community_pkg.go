package repositoryimpl

import (
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
)

type communityPkgImpl struct {
	appDB  dbimpl
	baseDB dbimpl
}

func (c communityPkgImpl) AddApplicationPkg(app domain.ApplicationPackage) error {
	res := c.toApplicationPkgDO(app)

	for i := range res {
		v := &res[i]

		if err := c.appDB.FirstOrCreate(
			nil,
			&applicationPkgDO{Community: v.Community, PackageName: v.PackageName, Version: v.Version, Repo: v.Repo}, v,
		); err != nil {
			return err
		}
	}

	return nil
}

func (c communityPkgImpl) AddBasePkg(app domain.BasePackage) error {
	res := c.toBasePkgDO(app)

	for i := range res {
		v := &res[i]

		if err := c.baseDB.FirstOrCreate(
			nil,
			&basePkgDO{Community: v.Community, PackageName: v.PackageName, Version: v.Version, Repo: v.Repo}, v,
		); err != nil {
			return err
		}
	}

	return nil
}
