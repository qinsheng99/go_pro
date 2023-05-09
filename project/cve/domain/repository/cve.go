package repository

import (
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
)

type CVE interface {
	FindOriginRecord(dp.CVENum) (domain.CveOriginRecordInfo, error)
	AddOriginRecord(*domain.CveOriginRecordInfo) error
	SaveOriginRecord(v *domain.CveOriginRecordInfo) error
}
