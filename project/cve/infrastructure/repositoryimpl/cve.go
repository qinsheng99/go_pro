package repositoryimpl

import (
	"github.com/qinsheng99/go-domain-web/common/infrastructure/postgresql"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
)

type cveImpl struct {
	originRecord
}

func NewCVEImpl(cfg *postgresql.Config) repository.CVE {
	return &cveImpl{
		originRecord: originRecord{
			postgresql.NewPgDao(cfg.Table.CveOriginRecord),
		},
	}
}
