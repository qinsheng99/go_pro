package repository

import (
	"github.com/qinsheng99/go-domain-web/api/osv_api"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
)

type ROeCompatibilityOsv struct {
	mysql.OeCompatibilityOsv
	Updateime      string           `json:"updateTime,omitempty"`
	ToolsResult    []osv_api.Record `gorm:"column:tools_result" json:"toolsResult,omitempty"`
	PlatformResult []osv_api.Record `gorm:"column:platform_result" json:"platformResult,omitempty"`
}

type RepoOsvImpl interface {
	SyncOsv() (string, error)
	Find() ([]ROeCompatibilityOsv, int64, error)
}
