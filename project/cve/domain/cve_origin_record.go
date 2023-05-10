package domain

import (
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
	"github.com/qinsheng99/go-domain-web/utils"
)

type CveOriginRecordInfo struct {
	Id        string
	CreatedAt int64

	Status dp.CVEStatus

	Source Source

	CveSourceData

	BaseOrigin
}

type BaseOrigin struct {
	Pushed    string
	PushType  string
	Published string
}

type CveSourceData struct {
	CVENum         dp.CVENum
	Desc           dp.Description
	Patch          []Patch
	Severity       []Severity
	ReferencesData []ReferencesData
	Affected       []dp.Purl
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

type ReferencesData struct {
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

func NewCveOriginRecordInfo(s dp.Source, base BaseOrigin, cve CveSourceData) CveOriginRecordInfo {
	return CveOriginRecordInfo{
		CreatedAt: utils.Now(),
		Source: Source{
			Source:        s,
			UpdatedSource: s,
		},
		CveSourceData: cve,
		BaseOrigin:    base,
		Status:        dp.Add,
	}
}
