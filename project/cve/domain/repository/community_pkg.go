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

type optToFindPkgs struct {
	Community dp.Community

	PageNum      int
	CountPerPage int
}

type OptToFindBasePkg struct {
	Community dp.Community
	Name      dp.PackageName
}

type OptFindBasePkgs = optToFindPkgs
type OptFindApplicationPkgs = optToFindPkgs

type optToDeletePkgs struct {
	// delete not equal UpdatedAt record
	UpdatedAt string

	Community dp.Community
}

type OptToDeleteApplicationPkgs = optToDeletePkgs
type OptToDeleteBasePkgs = optToDeletePkgs

type BasePkgRepository interface {
	AddBasePkg(*domain.BasePackage) error

	FindBasePkgs(OptFindBasePkgs) (v []domain.BasePackage, err error)
	FindBasePkg(OptToFindBasePkg) (domain.BasePackage, error)

	SaveBasePkg(*domain.BasePackage) error

	DeleteBasePkgs(OptToDeleteBasePkgs) error
}

type ApplicationPkgRepository interface {
	AddApplicationPkg(*domain.ApplicationPackage) error

	FindApplicationPkgs(OptFindApplicationPkgs) ([]domain.ApplicationPackage, error)
	FindApplicationPkg(OptToFindApplicationPkg) (domain.ApplicationPackage, error)

	SaveApplicationPkg(*domain.ApplicationPackage) error

	DeleteApplicationPkgs(OptToDeleteApplicationPkgs) error
}
