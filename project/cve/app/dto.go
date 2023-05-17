package app

import (
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
)

type CmdToAddCVEBasicInfo struct {
	Source dp.Source
	CVENum dp.CVENum

	CveApplication
}

type Severity = domain.Severity
type References = domain.References
type Patch = domain.Patch
type CveApplication = domain.CveApplication

type DetailInfoDTO struct {
	Desc       string       `json:"desc"`
	CveNum     string       `json:"cve_num"`
	Source     string       `json:"source"`
	Pushed     string       `json:"pushed"`
	PushType   string       `json:"push_type"`
	Affected   []string     `json:"affected"`
	Published  string       `json:"published"`
	Severity   []Severity   `json:"severity"`
	References []References `json:"references"`
	Patch      []Patch      `json:"patch"`
}

func toDetailInfoDTO(v *domain.CveBasicInfo) DetailInfoDTO {
	app := &v.CveApplication
	d := DetailInfoDTO{
		CveNum:     v.CVENum.CVENum(),
		Desc:       app.Desc.CveDescription(),
		Source:     v.Source.Source.Source(),
		Pushed:     app.Basic.Pushed,
		PushType:   app.Basic.PushType,
		Affected:   nil,
		Published:  app.Basic.Published,
		Severity:   app.Severity,
		References: app.References,
		Patch:      app.Patch,
	}

	d.Affected = make([]string, len(app.Affected))
	for i := range app.Affected {
		d.Affected[i] = app.Affected[i].Purl()
	}

	return d
}
