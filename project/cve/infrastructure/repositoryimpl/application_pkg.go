package repositoryimpl

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/dao"
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
)

func (a applicationPkgImpl) AddApplicationPkg(app *domain.ApplicationPackage) error {
	res := a.toApplicationPkgDO(app)

	for i := range res {
		v := &res[i]

		if err := a.db.Insert(nil, v); err != nil {
			return err
		}
	}

	return nil
}

func (a applicationPkgImpl) FindApplicationPkgs(opts repository.OptFindApplicationPkgs) (v []domain.ApplicationPackage, err error) {
	var do []applicationPkgDO
	err = a.db.GetRecords(
		a.db.DB(),
		func(db *gorm.DB) *gorm.DB {
			if opts.UpdatedAt != "" {
				db.Where("updated_at <> ?", opts.UpdatedAt)
			}

			if opts.Community != nil {
				db.Where("community = ?", opts.Community.Community())
			}

			return db
		},
		&do, dao.Pagination{}, nil,
	)
	if err != nil {
		return
	}

	f := func(community, repo string) int {
		for i := range v {
			r := &v[i].Repository
			if r.Repo == repo && r.Community != nil && r.Community.Community() == community {
				return i
			}
		}

		return -1
	}

	for i := range do {
		if appPkg, err := do[i].toApplicationPkg(); err != nil {
			return nil, err
		} else {
			if idx := f(appPkg.Repository.Community.Community(), appPkg.Repository.Repo); idx == -1 {
				v = append(v, appPkg)
			} else {
				v[idx].Packages = append(v[idx].Packages, appPkg.Packages...)
			}
		}
	}

	return
}

func (a applicationPkgImpl) FindApplicationPkg(opts repository.OptToFindApplicationPkg) (domain.ApplicationPackage, error) {
	var do applicationPkgDO
	err := a.db.GetRecord(
		a.db.DB(),
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

func (a applicationPkgImpl) DeleteApplicationPkg(id string) error {
	u, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return a.db.Delete(a.db.DB(), &applicationPkgDO{Id: u})
}