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

type basePkgImpl struct {
	db dbimpl
}

type applicationPkgImpl struct {
	db dbimpl
}

func NewApplicationPkgImpl(cfg *postgres.Config) repository.ApplicationPkgRepository {
	return &applicationPkgImpl{
		db: postgres.NewPgDao(cfg.Table.ApplicationPkg),
	}
}

func NewBasePkgImpl(cfg *postgres.Config) repository.BasePkgRepository {
	return &basePkgImpl{
		db: postgres.NewPgDao(cfg.Table.BasePkg),
	}
}
