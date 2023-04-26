package app

import (
	"encoding/json"

	"github.com/qinsheng99/go-domain-web/common/api"
	"github.com/qinsheng99/go-domain-web/domain"
	_const "github.com/qinsheng99/go-domain-web/utils/const"
)

type CompatibilityOsv struct {
	domain.CompatibilityOsv
	Updateime      string       `json:"updateTime,omitempty"`
	ToolsResult    []api.Record `gorm:"column:tools_result" json:"toolsResult,omitempty"`
	PlatformResult []api.Record `gorm:"column:platform_result" json:"platformResult,omitempty"`
}

type compatibilityOsvDTO struct {
	OsvList []CompatibilityOsv `json:"osv_list"`
	Total   int64              `json:"total"`
}

func toCompatibilityOsvDTO(list []domain.CompatibilityOsv, total int64) *compatibilityOsvDTO {
	data := make([]CompatibilityOsv, len(list))
	for i, v := range list {
		var t []api.Record
		_ = json.Unmarshal([]byte(v.ToolsResult), &t)
		var p []api.Record
		_ = json.Unmarshal([]byte(v.PlatformResult), &p)

		data[i] = CompatibilityOsv{
			CompatibilityOsv: v,
			ToolsResult:      t,
			PlatformResult:   p,
			Updateime:        v.Updateime.Format(_const.Format),
		}
	}

	return &compatibilityOsvDTO{OsvList: data, Total: total}
}
