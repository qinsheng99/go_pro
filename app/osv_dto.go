package app

import (
	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
)

type ROeCompatibilityOsv struct {
	mysql.OeCompatibilityOsv
	Updateime      string       `json:"updateTime,omitempty"`
	ToolsResult    []api.Record `gorm:"column:tools_result" json:"toolsResult,omitempty"`
	PlatformResult []api.Record `gorm:"column:platform_result" json:"platformResult,omitempty"`
}

type ResultOsv struct {
	OsvList []ROeCompatibilityOsv `json:"osv_list"`
	Total   int64                 `json:"total"`
}
