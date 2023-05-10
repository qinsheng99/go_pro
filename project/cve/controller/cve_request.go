package controller

import (
	"errors"

	"github.com/qinsheng99/go-domain-web/project/cve/app"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
)

type uvpDataRequest struct {
	Id         string               `json:"id"         binding:"required"`
	Desc       string               `json:"desc"       binding:"required"`
	Source     string               `json:"source"     binding:"required"`
	Pushed     string               `json:"pushed"     binding:"required"`
	PushType   string               `json:"push_type"  binding:"required"`
	Affected   []string             `json:"affected"   binding:"required"`
	Published  string               `json:"published"  binding:"required"`
	Severity   []app.Severity       `json:"severity"`
	References []app.ReferencesData `json:"references"`
	Patch      []app.Patch          `json:"patch"`
}

func (u *uvpDataRequest) toCmd() (v app.OriginRecordCmd, err error) {
	v.BaseOrigin = app.BaseOrigin{
		Pushed:    u.Pushed,
		PushType:  u.PushType,
		Published: u.Published,
	}

	s := &v.CveSourceData
	for _, a := range u.Affected {
		if p, err := dp.NewPurl(a); err == nil {
			s.Affected = append(v.Affected, p)
		}
	}

	if len(s.Affected) == 0 {
		err = errors.New("affected is empty")

		return
	}

	s.Patch = u.Patch
	s.Severity = u.Severity
	s.ReferencesData = u.References
	if v.Source, err = dp.NewSource(u.Source); err != nil {
		return
	}

	s.Desc = dp.NewDescription(u.Desc)
	s.CVENum, err = dp.NewCVENum(u.Id)

	return
}
