package repository

import (
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
)

type OptToFindApplicationPkg struct {
	Repo    string
	Version string

	Community dp.Community
	Name      dp.PackageName
}

type OptToFindPkgs struct {
	Community dp.Community

	PageNum      int
	CountPerPage int
}

type OptToFindBasePkg struct {
	Version string

	Community dp.Community
	Name      dp.PackageName
}

type PkgImpl interface {
	AddApplicationPkg(domain.ApplicationPackage) error
	AddBasePkg(domain.BasePackage) error

	FindApplicationPkgs(dp.Community, int64) ([]domain.ApplicationPackage, error)
	FindApplicationPkg(OptToFindApplicationPkg) (domain.ApplicationPackage, error)
	DeleteApplicationPkg(id string) error

	FindBasePkgs(OptToFindPkgs, int64) (v []domain.BasePackage, err error)
	FindBasePkg(OptToFindBasePkg) (domain.BasePackage, error)
	DeleteBasePkg(id string) error
}
