package repository

import (
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
)

type CVE interface {
	FindCVEBasicInfo(dp.CVENum) (domain.CveBasicInfo, error)
	AddCVEBasicInfo(*domain.CveBasicInfo) error
	SaveCVEBasicInfo(v *domain.CveBasicInfo) error
}
