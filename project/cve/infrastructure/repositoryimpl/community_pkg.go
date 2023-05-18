package repositoryimpl

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/dao"
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
)

type communityPkgImpl struct {
	appDB  dbimpl
	baseDB dbimpl
}

func (c communityPkgImpl) AddApplicationPkg(app *domain.ApplicationPackage) error {
	res := c.toApplicationPkgDO(app)

	for i := range res {
		v := &res[i]

		if err := c.appDB.Insert(nil, v); err != nil {
			return err
		}
	}

	return nil
}

func (c communityPkgImpl) AddBasePkg(app *domain.BasePackage) error {
	var do basePkgDO
	c.toBasePkgDO(app, &do)

	return c.baseDB.Insert(nil, &do)
}

func (c communityPkgImpl) FindApplicationPkgs(opts repository.OptToFindPkgs) (v []domain.ApplicationPackage, err error) {
	var do []applicationPkgDO
	err = c.appDB.GetRecords(
		c.appDB.DB(),
		func(db *gorm.DB) *gorm.DB {
			if len(opts.UpdatedAt) != 0 {
				db.Where("updated_at <> ?", opts.UpdatedAt)
			}

			return db.Where("community = ?", opts.Community.Community())
		},
		&do, dao.Pagination{}, nil,
	)
	if err != nil {
		return
	}

	f := func(repo string) int {
		for i := range v {
			if v[i].Repository.Repo == repo {
				return i
			}
		}

		return -1
	}

	for i := range do {
		if appPkg, err := do[i].toApplicationPkg(); err != nil {
			return nil, err
		} else {
			if idx := f(appPkg.Repository.Repo); idx == -1 {
				v = append(v, appPkg)
			} else {
				v[idx].Packages = append(v[idx].Packages, appPkg.Packages...)
			}
		}
	}

	return
}

func (c communityPkgImpl) FindApplicationPkg(opts repository.OptToFindApplicationPkg) (domain.ApplicationPackage, error) {
	var do applicationPkgDO
	err := c.appDB.GetRecord(
		c.appDB.DB(),
		func(db *gorm.DB) *gorm.DB {
			return db.Where(
				"community = ? and repo = ? and package_name = ? and version = ?",
				opts.Community.Community(), opts.Repo, opts.Name.PackageName(), opts.Version,
			)
		},
		&do,
	)

	if err != nil {
		return domain.ApplicationPackage{}, err
	}

	return do.toApplicationPkg()
}

func (c communityPkgImpl) FindBasePkgs(opts repository.OptToFindPkgs) (v []domain.BasePackage, err error) {
	var do []basePkgDO
	err = c.baseDB.GetRecords(
		c.baseDB.DB(),
		func(db *gorm.DB) *gorm.DB {
			if len(opts.UpdatedAt) != 0 {
				db.Where("updated_at <> ?", opts.UpdatedAt)
			}

			return db.Where("community = ?", opts.Community.Community())
		},
		&do, dao.Pagination{
			PageNum:      opts.PageNum,
			CountPerPage: opts.CountPerPage,
		}, []dao.SortByColumn{{Column: "created_at"}},
	)
	if err != nil {
		return
	}

	v = make([]domain.BasePackage, len(do))

	for i := range do {
		if v[i], err = do[i].toBasePkg(); err != nil {
			return
		}
	}

	return
}

func (c communityPkgImpl) FindBasePkg(opts repository.OptToFindBasePkg) (domain.BasePackage, error) {
	var do basePkgDO

	err := c.baseDB.GetRecord(c.baseDB.DB(), func(db *gorm.DB) *gorm.DB {
		return db.Where(
			"community = ? and package_name = ?",
			opts.Community.Community(), opts.Name.PackageName(),
		)
	}, &do)

	if err != nil {
		return domain.BasePackage{}, err
	}

	return do.toBasePkg()
}

func (c communityPkgImpl) DeleteApplicationPkg(id string) error {
	u, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return c.appDB.Delete(c.appDB.DB(), &applicationPkgDO{Id: u})
}

func (c communityPkgImpl) DeleteBasePkg(id string) error {
	u, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return c.baseDB.Delete(c.baseDB.DB(), &basePkgDO{Id: u})
}
