package app

import (
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
)

type CveOriginService interface {
}

type cveOriginService struct {
	repo repository.CVE
}

func NewCveOriginService(repo repository.CVE) CveOriginService {
	return &cveOriginService{
		repo: repo,
	}
}

func (c *cveOriginService) AddCVEOriginRecord(app *OriginRecordCmd) error {
	return nil
}

func (c *cveOriginService) findCVEOriginRecord(v dp.CVENum) (domain.CveOriginRecordInfo, error) {
	return domain.CveOriginRecordInfo{}, nil
}

func (c *cveOriginService) saveCVEOriginRecord(v *domain.CveOriginRecordInfo) error {
	return nil
}
