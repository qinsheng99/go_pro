package repositoryimpl

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"

	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
	"github.com/qinsheng99/go-domain-web/utils"
)

type cveBasicInfoDO struct {
	Id            uuid.UUID      `gorm:"column:uuid;type:uuid"                        json:"-"`
	Desc          string         `gorm:"column:desc"                                  json:"desc"`
	Source        string         `gorm:"column:source"`
	CveNum        string         `gorm:"column:cve_num"`
	Pushed        string         `gorm:"column:pushed"`
	Status        string         `gorm:"column:cve_status"`
	PushType      string         `gorm:"column:push_type"`
	Published     string         `gorm:"column:published" `
	UpdatedSource string         `gorm:"column:updated_source"`
	CreatedAt     int64          `gorm:"column:created_at"`
	UpdatedAt     int64          `gorm:"column:updated_at"`
	Affected      pq.StringArray `gorm:"column:affected;type:text[];default:'{}'"`
	Patch         postgres.Jsonb `gorm:"column:patch;type:jsonb;default:'{}'"`
	Severity      postgres.Jsonb `gorm:"column:severity;type:jsonb;default:'{}'"`
	References    postgres.Jsonb `gorm:"column:references;type:jsonb;default:'{}'"`
}

func (do *cveBasicInfoDO) toCveOriginRecordInfo() (v domain.CveBasicInfo, err error) {
	v.Id = do.Id.String()

	if v.CVENum, err = dp.NewCVENum(do.CveNum); err != nil {
		return
	}
	if v.Source.Source, err = dp.NewSource(do.Source); err != nil {
		return
	}

	if v.Source.UpdatedSource, err = dp.NewSource(do.UpdatedSource); err != nil {
		return
	}

	app := &v.CveApplication
	app.Basic.PushType = do.PushType
	app.Basic.Published = do.Published
	app.Basic.CreatedAt = do.CreatedAt
	app.Basic.Pushed = do.Pushed

	app.Desc = dp.NewDescription(do.Desc)

	if app.Basic.Status, err = dp.NewCVEStatus(do.Status); err != nil {
		return
	}

	app.Affected = make([]dp.Purl, len(do.Affected))

	for i := range do.Affected {
		if app.Affected[i], err = dp.NewPurl(do.Affected[i]); err != nil {
			return
		}
	}

	if err = json.Unmarshal(do.Patch.RawMessage, &app.Patch); err != nil {
		return
	}

	if err = json.Unmarshal(do.Severity.RawMessage, &app.Severity); err != nil {
		return
	}

	err = json.Unmarshal(do.References.RawMessage, &app.References)

	return
}

func (o originRecord) toCveBasicInfoDO(v *domain.CveBasicInfo) (do cveBasicInfoDO, err error) {
	app := &v.CveApplication
	do = cveBasicInfoDO{
		Desc:          v.Desc.Description(),
		Source:        v.Source.Source.Source(),
		CveNum:        v.CVENum.CVENum(),
		Pushed:        app.Basic.Pushed,
		Status:        app.Basic.Status.CVEStatus(),
		PushType:      app.Basic.PushType,
		Published:     app.Basic.Published,
		UpdatedSource: v.Source.UpdatedSource.Source(),
		CreatedAt:     app.Basic.CreatedAt,
		UpdatedAt:     app.Basic.CreatedAt,
	}

	do.Affected = make(pq.StringArray, len(app.Affected))

	for i := range app.Affected {
		do.Affected[i] = app.Affected[i].Purl()
	}

	if do.Patch, err = utils.ToJsonB(app.Patch); err != nil {
		return
	}

	if do.Severity, err = utils.ToJsonB(app.Severity); err != nil {
		return
	}

	if do.References, err = utils.ToJsonB(app.References); err != nil {
		return
	}

	return
}
