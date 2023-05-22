package repositoryimpl

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/dao"
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
)

func (b basePkgImpl) AddBasePkg(app *domain.BasePackage) error {
	var do basePkgDO
	b.toBasePkgDO(app, &do)

	return b.db.Insert(nil, &do)
}

func (b basePkgImpl) FindBasePkgs(opts repository.OptFindBasePkgs) (v []domain.BasePackage, err error) {
	var do []basePkgDO
	err = b.db.GetRecords(
		b.db.DB(),
		func(db *gorm.DB) *gorm.DB {
			if opts.Community != nil {
				db.Where("community = ?", opts.Community.Community())
			}

			return db
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

func (b basePkgImpl) FindBasePkg(opts repository.OptToFindBasePkg) (domain.BasePackage, error) {
	var do basePkgDO

	err := b.db.GetRecord(b.db.DB(), func(db *gorm.DB) *gorm.DB {
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

func (b basePkgImpl) DeleteBasePkgs(up string) error {
	return b.db.Delete(b.db.DB(), func(db *gorm.DB) *gorm.DB {
		return db.Where(updatedAt+"<> ?", up)
	})
}

func (b basePkgImpl) SaveBasePkg(app *domain.BasePackage) error {
	var do basePkgDO
	b.toBasePkgDO(app, &do)

	id, err := uuid.Parse(app.Id)
	if err != nil {
		return err
	}

	return b.db.UpdateRecord(b.db.DB(), &basePkgDO{Id: id}, do.toUpdatedMap())
}
