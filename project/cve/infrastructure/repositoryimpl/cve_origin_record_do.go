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

type cveOriginRecordDO struct {
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

func (do *cveOriginRecordDO) toCveOriginRecordInfo() (v domain.CveOriginRecordInfo, err error) {
	v.Id = do.Id.String()
	v.PushType = do.PushType
	v.Published = do.Published
	v.CreatedAt = do.CreatedAt
	v.Pushed = do.Pushed

	v.Affected = make([]dp.Purl, len(do.Affected))

	for i := range do.Affected {
		if v.Affected[i], err = dp.NewPurl(do.Affected[i]); err != nil {
			return
		}
	}

	if err = json.Unmarshal(do.Patch.RawMessage, &v.Patch); err != nil {
		return
	}

	if v.Source.Source, err = dp.NewSource(do.Source); err != nil {
		return
	}

	if v.Source.UpdatedSource, err = dp.NewSource(do.UpdatedSource); err != nil {
		return
	}

	if err = json.Unmarshal(do.Severity.RawMessage, &v.Severity); err != nil {
		return
	}

	if err = json.Unmarshal(do.References.RawMessage, &v.ReferencesData); err != nil {
		return
	}

	if v.CVENum, err = dp.NewCVENum(do.CveNum); err != nil {
		return
	}

	v.Desc = dp.NewDescription(do.Desc)

	v.Status, err = dp.NewCVEStatus(do.Status)

	return
}

func (o originRecord) toCveOriginRecordDO(v *domain.CveOriginRecordInfo) (do cveOriginRecordDO, err error) {
	do = cveOriginRecordDO{
		Desc:          v.Desc.Description(),
		Source:        v.Source.Source.Source(),
		CveNum:        v.CVENum.CVENum(),
		Pushed:        v.Pushed,
		Status:        v.Status.CVEStatus(),
		PushType:      v.PushType,
		Published:     v.Published,
		UpdatedSource: v.Source.UpdatedSource.Source(),
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     v.CreatedAt,
	}

	do.Affected = make(pq.StringArray, len(v.Affected))

	for i := range v.Affected {
		do.Affected[i] = v.Affected[i].Purl()
	}

	if do.Patch, err = utils.ToJsonB(v.Patch); err != nil {
		return
	}

	if do.Severity, err = utils.ToJsonB(v.Severity); err != nil {
		return
	}

	if do.References, err = utils.ToJsonB(v.ReferencesData); err != nil {
		return
	}

	return
}
