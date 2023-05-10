package repositoryimpl

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
)

type originRecord struct {
	cli dbimpl
}

func (o originRecord) FindCVEBasicInfo(num dp.CVENum) (domain.CveBasicInfo, error) {
	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("cve_num = ?", num.CVENum())
	}

	var res cveBasicInfoDO

	if err := o.cli.GetRecord(filter, &res); err != nil {
		return domain.CveBasicInfo{}, err
	}

	return res.toCveOriginRecordInfo()
}

func (o originRecord) AddCVEBasicInfo(v *domain.CveBasicInfo) error {
	do, err := o.toCveBasicInfoDO(v)
	if err != nil {
		return err
	}

	err = o.cli.Insert(&cveBasicInfoDO{CveNum: do.CveNum}, &do)
	if err != nil {
		return err
	}

	v.Id = do.Id.String()

	return err
}

func (o originRecord) SaveCVEBasicInfo(v *domain.CveBasicInfo) error {
	do, err := o.toCveBasicInfoDO(v)
	if err != nil {
		return err
	}

	u, err := uuid.Parse(v.Id)
	if err != nil {
		return err
	}

	return o.cli.UpdateRecord(&cveBasicInfoDO{Id: u}, &do)
}
