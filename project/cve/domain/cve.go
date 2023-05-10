package domain

import (
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
	"github.com/qinsheng99/go-domain-web/utils"
)

type CveBasicInfo struct {
	Id     string
	CVENum dp.CVENum
	Source Source

	CveApplication
}

type CveRecord struct {
	Pushed    string
	PushType  string
	Published string
	CreatedAt int64

	Status dp.CVEStatus
}

type CveApplication struct {
	Basic CveRecord

	Desc       dp.Description
	Patch      []Patch
	Affected   []dp.Purl
	Severity   []Severity
	References []References
}

type Source struct {
	Source        dp.Source
	UpdatedSource dp.Source
}

type Severity struct {
	Type   string `json:"type"`
	Score  string `json:"score"`
	Vector string `json:"vector"`
}

type References struct {
	Url  string `json:"url"`
	Type string `json:"type"`
}

type Patch struct {
	Package    string `json:"package"`
	FixVersion string `json:"fix_version"`
	FixPatch   string `json:"fix_patch"`
	BreakPatch string `json:"break_patch"`
	Source     string `json:"source"`
	Branch     string `json:"branch"`
}

func (entity *CveBasicInfo) UpdateStatus(v dp.CVEStatus) {
	entity.CveApplication.Basic.Status = v
}

func (entity *CveBasicInfo) UpdateCveApplication(v *CveApplication) {
	entity.CveApplication = *v
}

func NewCveBasicInfo(s dp.Source, app CveApplication, cve dp.CVENum) CveBasicInfo {
	app.Basic.CreatedAt = utils.Now()
	app.Basic.Status = dp.Add
	return CveBasicInfo{
		Source: Source{
			Source:        s,
			UpdatedSource: s,
		},
		CVENum:         cve,
		CveApplication: app,
	}
}
