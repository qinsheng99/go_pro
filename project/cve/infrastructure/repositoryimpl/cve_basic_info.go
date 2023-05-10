package repositoryimpl

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
	"github.com/qinsheng99/go-domain-web/utils"
)

type basicInfo struct {
	cli dbimpl
}

func (o basicInfo) FindCVEBasicInfo(num dp.CVENum) (domain.CveBasicInfo, error) {
	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("cve_num = ?", num.CVENum())
	}

	var res cveBasicInfoDO

	if err := o.cli.GetRecord(filter, &res); err != nil {
		return domain.CveBasicInfo{}, err
	}

	return res.toCveBasicInfo()
}

func (o basicInfo) AddCVEBasicInfo(v *domain.CveBasicInfo) error {
	do, err := o.toCveBasicInfoDO(v)
	if err != nil {
		return err
	}

	res, err := do.toMap()
	if err != nil {
		return err
	}

	res[fieldId] = do.Id.String()
	res[fieldCreatedAt] = v.Basic.CreatedAt
	res[fieldCveNum] = do.CveNum
	res[fieldSource] = do.Source

	err = o.cli.Insert(&cveBasicInfoDO{CveNum: do.CveNum}, res)
	if err != nil {
		return err
	}

	v.Id = do.Id.String()

	return err
}

func (o basicInfo) SaveCVEBasicInfo(v *domain.CveBasicInfo) error {
	do, err := o.toCveBasicInfoDO(v)
	if err != nil {
		return err
	}

	do.UpdatedAt = utils.Now()

	d, err := do.toMap()
	if err != nil {
		return err
	}

	u, err := uuid.Parse(v.Id)
	if err != nil {
		return err
	}

	return o.cli.UpdateRecord(&cveBasicInfoDO{Id: u}, d)
}
