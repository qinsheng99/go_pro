package app

import (
	"encoding/json"

	"github.com/qinsheng99/go-domain-web/common/api"
	"github.com/qinsheng99/go-domain-web/domain"
	_const "github.com/qinsheng99/go-domain-web/utils/const"
)

type CompatibilityOsv struct {
	domain.CompatibilityOsvInfo
	Updateime      string       `json:"updateTime,omitempty"`
	ToolsResult    []api.Record `gorm:"column:tools_result" json:"toolsResult,omitempty"`
	PlatformResult []api.Record `gorm:"column:platform_result" json:"platformResult,omitempty"`
}

type CompatibilityOsvDTO struct {
	OsvList []CompatibilityOsv `json:"list"`
	Total   int64              `json:"total"`
}

func toCompatibilityOsvDTO(list []domain.CompatibilityOsvInfo, total int64) *CompatibilityOsvDTO {
	data := make([]CompatibilityOsv, len(list))
	for i, v := range list {
		var t []api.Record
		_ = json.Unmarshal([]byte(v.ToolsResult), &t)
		var p []api.Record
		_ = json.Unmarshal([]byte(v.PlatformResult), &p)

		data[i] = CompatibilityOsv{
			CompatibilityOsvInfo: v,
			ToolsResult:          t,
			PlatformResult:       p,
			Updateime:            v.Updateime.Format(_const.Format),
		}
	}

	return &CompatibilityOsvDTO{OsvList: data, Total: total}
}
