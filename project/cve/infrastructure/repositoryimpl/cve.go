package repositoryimpl

import (
	"github.com/qinsheng99/go-domain-web/common/infrastructure/postgres"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
)

type cveImpl struct {
	basicInfo
	communityPkg
}

func NewCVEImpl(cfg *postgres.Config) repository.CVE {
	return &cveImpl{
		basicInfo: basicInfo{
			postgres.NewPgDao(cfg.Table.CveBasicInfo),
		},
		communityPkg: communityPkg{
			postgres.NewPgDao(cfg.Table.CveCommunityPkg),
		},
	}
}
