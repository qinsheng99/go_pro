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
