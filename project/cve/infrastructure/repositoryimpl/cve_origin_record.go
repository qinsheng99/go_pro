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

func (o originRecord) FindOriginRecord(num dp.CVENum) (domain.CveOriginRecordInfo, error) {
	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("cve_num = ?", num.CVENum())
	}

	var res cveOriginRecordDO

	if err := o.cli.GetRecord(filter, &res); err != nil {
		return domain.CveOriginRecordInfo{}, err
	}

	return res.toCveOriginRecordInfo()
}

func (o originRecord) AddOriginRecord(v *domain.CveOriginRecordInfo) error {
	do, err := o.toCveOriginRecordDO(v)
	if err != nil {
		return err
	}

	err = o.cli.Insert(&cveOriginRecordDO{CveNum: do.CveNum}, &do)
	if err != nil {
		return err
	}

	v.Id = do.Id.String()

	return err
}

func (o originRecord) SaveOriginRecord(v *domain.CveOriginRecordInfo) error {
	do, err := o.toCveOriginRecordDO(v)
	if err != nil {
		return err
	}

	u, err := uuid.Parse(v.Id)
	if err != nil {
		return err
	}

	return o.cli.UpdateRecord(&cveOriginRecordDO{Id: u}, &do)
}
