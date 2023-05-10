package app

import (
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
)

type CveService interface {
	AddCVEBasicInfo(app *CmdToAddCVEBasicInfo) error
	BasicInfo(dp.CVENum) (DetailInfoDTO, error)
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

	_, err := c.repo.FindCVEBasicInfo(app.CVENum)
	if err == nil {
		return c.updateCve(app)
	}

	err = c.repo.AddCVEBasicInfo(&info)
	if err != nil {
		return err
	}

	// TODO send message  info.CVENum.CVENum

	return nil
}

func (c *cveService) BasicInfo(cve dp.CVENum) (DetailInfoDTO, error) {
	info, err := c.repo.FindCVEBasicInfo(cve)
	if err != nil {
		return DetailInfoDTO{}, err
	}

	return toDetailInfoDTO(&info), err
}

func (c *cveService) updateCve(app *CmdToAddCVEBasicInfo) error {
	res, err := c.repo.FindCVEBasicInfo(app.CVENum)
	if err != nil {
		return err
	}

	res.UpdateCveApplication(&app.CveApplication)
	res.UpdateStatus(dp.Update)

	err = c.repo.SaveCVEBasicInfo(&res)
	if err != nil {
		return err
	}

	// TODO send message  res.CVENum.CVENum

	return nil
}
