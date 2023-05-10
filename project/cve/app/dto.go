package app

import (
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
)

type OriginRecordCmd struct {
	Source dp.Source

	BaseOrigin
	CveSourceData
}

type Severity = domain.Severity
type ReferencesData = domain.ReferencesData
type Patch = domain.Patch
type BaseOrigin = domain.BaseOrigin
type CveSourceData = domain.CveSourceData
