package app

import (
	"github.com/sirupsen/logrus"

	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
)

type PkgService interface {
	AddApplicationPkg(*CmdToApplicationPkg) error
	AddBasePkg(*CmdToBasePkg) error

	ListBasePkgs(repository.OptFindBasePkgs) ([]ListBasePkgsDTO, error)
	ListApplicationPkgs(repository.OptFindApplicationPkgs) ([]ListApplicationPkgsDTO, error)
}

type pkgService struct {
	base        repository.BasePkgRepository
	application repository.ApplicationPkgRepository
}

func NewPkgService(base repository.BasePkgRepository, application repository.ApplicationPkgRepository) PkgService {
	return &pkgService{
		base:        base,
		application: application,
	}
}

func (p *pkgService) AddApplicationPkg(pkg *CmdToApplicationPkg) error {
	err := p.application.AddApplicationPkg(pkg)
	if err != nil {
		logrus.Errorf(
			"add application failed, community:%s, err:%s", pkg.Repository.Community.Community(), err.Error(),
		)
	}

	return nil
}

func (p *pkgService) AddBasePkg(pkg *CmdToBasePkg) error {
	err := p.base.AddBasePkg(pkg)
	if err != nil {
		logrus.Errorf(
			"add application failed, community:%s, err:%s", pkg.Repository.Community.Community(), err.Error(),
		)
	}

	return nil
}

func (p *pkgService) ListBasePkgs(opts repository.OptFindBasePkgs) ([]ListBasePkgsDTO, error) {
	pkgs, err := p.base.FindBasePkgs(opts)
	if err != nil {
		return nil, err
	}

	return toListBasePkgsDTO(pkgs)
}

func (p *pkgService) ListApplicationPkgs(opts repository.OptFindApplicationPkgs) ([]ListApplicationPkgsDTO, error) {
	pkgs, err := p.application.FindApplicationPkgs(opts)
	if err != nil {
		return nil, err
	}

	return toListApplicationDTO(pkgs), nil
}
