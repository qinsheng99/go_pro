package domain

import (
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
)

type CveOriginRecordInfo struct {
	Id        string
	Pushed    string
	PushType  string
	Published string
	CreatedAt int64
	Affected  []string

	Patch          []Patch
	Source         Source
	Severity       []Severity
	ReferencesData []ReferencesData

	CVENum dp.CVENum
	Status dp.CVEStatus
	Desc   dp.Description
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
