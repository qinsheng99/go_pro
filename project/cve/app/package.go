package app

import (
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
)

type PkgService interface {
	AddApplicationPkg(pkg *CmdToApplicationPkg) error
}

type pkgService struct {
	repo repository.PkgImpl
}

func NewPkgService(repo repository.PkgImpl) PkgService {
	return &pkgService{
		repo: repo,
	}
}

func (p *pkgService) AddApplicationPkg(app *CmdToApplicationPkg) error {
	//err := p.repo.AddApplicationPkg(app)
	//if err != nil {
	//	logrus.Errorf(
	//		"add application failed, community:%s, err:%s", app.Community.Community(), err.Error(),
	//	)
	//}

	return nil
}
