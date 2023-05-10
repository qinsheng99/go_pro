package app

import (
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
)

type CveService interface {
	AddCVEBasicInfo(app *CmdToAddCVEBasicInfo) error
}

type cveService struct {
	repo repository.CVE
}

func NewCveService(repo repository.CVE) CveService {
	return &cveService{
		repo: repo,
	}
}

func (c *cveService) AddCVEBasicInfo(app *CmdToAddCVEBasicInfo) error {
	info := domain.NewCveBasicInfo(app.Source, app.CveApplication, app.CVENum)

	err := c.repo.AddCVEBasicInfo(&info)
	if err != nil {
		return err
	}

	// TODO send message  info.Id

	return nil
}
