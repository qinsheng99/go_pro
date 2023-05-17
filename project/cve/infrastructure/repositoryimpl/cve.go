package repositoryimpl

import (
	"github.com/qinsheng99/go-domain-web/common/infrastructure/postgres"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
)

type cveImpl struct {
	basicInfo
}

func NewCVEImpl(cfg *postgres.Config) repository.CVE {
	return &cveImpl{
		basicInfo: basicInfo{
			postgres.NewPgDao(cfg.Table.BasicInfo),
		},
	}
}

type pkgImpl struct {
	communityPkg
}

func NewPkgImpl(cfg *postgres.Config) repository.PkgImpl {
	return &pkgImpl{
		communityPkg: communityPkg{
			postgres.NewPgDao(cfg.Table.CommunityPkg),
		},
	}
}
