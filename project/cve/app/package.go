package app

import (
	"github.com/sirupsen/logrus"

	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
)

type PkgService interface {
	AddApplicationPkg(CmdToApplicationPkg) error
	AddBasePkg(CmdToBasePkg) error
}

type pkgService struct {
	repo repository.PkgImpl
}

func NewPkgService(repo repository.PkgImpl) PkgService {
	return &pkgService{
		repo: repo,
	}
}

func (p *pkgService) AddApplicationPkg(pkg CmdToApplicationPkg) error {
	err := p.repo.AddApplicationPkg(pkg)
	if err != nil {
		logrus.Errorf(
			"add application failed, community:%s, err:%s", pkg.Repository.Community.Community(), err.Error(),
		)
	}

	return nil
}

func (p *pkgService) AddBasePkg(pkg CmdToBasePkg) error {
	err := p.repo.AddBasePkg(pkg)
	if err != nil {
		logrus.Errorf(
			"add application failed, community:%s, err:%s", pkg.Repository.Community.Community(), err.Error(),
		)
	}

	return nil
}
