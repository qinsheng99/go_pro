package repositoryimpl

import (
	"github.com/qinsheng99/go-domain-web/common/infrastructure/postgresql"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
)

type cveImpl struct {
	basicInfo
}

func NewCVEImpl(cfg *postgresql.Config) repository.CVE {
	return &cveImpl{
		basicInfo: basicInfo{
			postgresql.NewPgDao(cfg.Table.CveBasicInfo),
		},
	}
}
