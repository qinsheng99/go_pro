package repositoryimpl

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"

	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
)

const (
	fieldId         = "uuid"
	fieldDesc       = "desc"
	fieldPatch      = "patch"
	fieldSource     = "source"
	fieldCveNum     = "cve_num"
	fieldAffected   = "affected"
	fieldSeverity   = "severity"
	fieldCreatedAt  = "created_at"
	fieldReferences = "references"
)

var updates = []string{
	"desc", "pushed", "cve_status", "push_type", "published", "updated_source", "affected", "severity", "references", "updated_at",
}

type cveBasicInfoDO struct {
	Id            uuid.UUID      `gorm:"column:uuid;type:uuid"                        json:"-"`
	Desc          string         `gorm:"column:desc"                                  json:"desc"`
	Source        string         `gorm:"column:source"                                json:"-"`
	CveNum        string         `gorm:"column:cve_num"                               json:"-"`
	Pushed        string         `gorm:"column:pushed"                                json:"pushed"`
	Status        string         `gorm:"column:cve_status"                            json:"cve_status"`
	PushType      string         `gorm:"column:push_type"                             json:"push_type"`
	Published     string         `gorm:"column:published"                             json:"published"`
	UpdatedSource string         `gorm:"column:updated_source"                        json:"updated_source"`
	Patch         string         `gorm:"column:patch;type:jsonb;default:'{}'"         json:"affected"`
	Severity      string         `gorm:"column:severity;type:jsonb;default:'{}'"      json:"severity"`
	References    string         `gorm:"column:references;type:jsonb;default:'{}'"    json:"references"`
	CreatedAt     int64          `gorm:"column:created_at"                            json:"-"`
	UpdatedAt     int64          `gorm:"column:updated_at"                            json:"updated_at"`
	Affected      pq.StringArray `gorm:"column:affected;type:text[];default:'{}'"     json:"-"`
}

func (do *cveBasicInfoDO) toMap() (res map[string]any, _ error) {
	v, err := json.Marshal(do)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(v, &res)

	res[fieldAffected] = marshalStringArray(do.Affected)

	return res, err
}

func marshalJsonb(v postgres.Jsonb) string {
	value, err := v.Value()
	if err != nil {
		return "{}"
	}

	if value != nil {
		if s, ok := value.([]byte); ok {
			return string(s)
		}
	}

	return "{}"
}

func marshalStringArray(sa pq.StringArray) string {
	v, err := sa.Value()
	if err != nil {
		return ""
	}

	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}

	return "{}"
}

func (do *cveBasicInfoDO) toCveBasicInfo() (v domain.CveBasicInfo, err error) {
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

	if err = json.Unmarshal([]byte(do.Patch), &app.Patch); err != nil {
		return
	}

	if err = json.Unmarshal([]byte(do.Severity), &app.Severity); err != nil {
		return
	}

	err = json.Unmarshal([]byte(do.References), &app.References)

	return
}

func (o basicInfo) toCveBasicInfoDO(v *domain.CveBasicInfo) (do cveBasicInfoDO, err error) {
	app := &v.CveApplication
	do = cveBasicInfoDO{
		Id:            uuid.New(),
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

	if do.Patch, err = toStr(app.Patch); err != nil {
		return
	}

	if do.Severity, err = toStr(app.Severity); err != nil {
		return
	}

	if do.References, err = toStr(app.References); err != nil {
		return
	}

	return
}

func toStr(v interface{}) (string, error) {
	bys, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(bys), err
}
