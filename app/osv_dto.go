package app

import (
	"encoding/json"

	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/domain"
	_const "github.com/qinsheng99/go-domain-web/utils/const"
)

type oeCompatibilityOsv struct {
	domain.OeCompatibilityOsv
	Updateime      string       `json:"updateTime,omitempty"`
	ToolsResult    []api.Record `gorm:"column:tools_result" json:"toolsResult,omitempty"`
	PlatformResult []api.Record `gorm:"column:platform_result" json:"platformResult,omitempty"`
}

type resultOsvDTO struct {
	OsvList []oeCompatibilityOsv `json:"osv_list"`
	Total   int64                `json:"total"`
}

func toResultOsvDTO(list []domain.OeCompatibilityOsv, total int64) *resultOsvDTO {
	data := make([]oeCompatibilityOsv, len(list))
	for i, v := range list {
		var t []api.Record
		_ = json.Unmarshal([]byte(v.ToolsResult), &t)
		var p []api.Record
		_ = json.Unmarshal([]byte(v.PlatformResult), &p)

		data[i] = oeCompatibilityOsv{
			OeCompatibilityOsv: v,
			ToolsResult:        t,
			PlatformResult:     p,
			Updateime:          v.Updateime.Format(_const.Format),
		}
	}
	return &resultOsvDTO{OsvList: data, Total: total}
}
