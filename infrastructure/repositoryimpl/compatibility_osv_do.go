package repositoryimpl

import (
	"time"

	"github.com/qinsheng99/go-domain-web/common/api"
	"github.com/qinsheng99/go-domain-web/domain"
)

type compatibilityOsvDO struct {
	Id                   int64     `gorm:"column:id"`
	Architecture         string    `gorm:"column:architecture"`
	OsVersion            string    `gorm:"column:os_version"`
	OsvName              string    `gorm:"column:osv_name"`
	Date                 string    `gorm:"column:date"`
	OsDownloadLink       string    `gorm:"column:os_download_link"`
	Type                 string    `gorm:"column:type"`
	Details              string    `gorm:"column:details"`
	FriendlyLink         string    `gorm:"column:friendly_link"`
	TotalResult          string    `gorm:"column:total_result"`
	CheckSum             string    `gorm:"column:checksum"`
	BaseOpeneulerVersion string    `gorm:"column:base_openeuler_version"`
	ToolsResult          string    `gorm:"column:tools_result"`
	PlatformResult       string    `gorm:"column:platform_result"`
	Updateime            time.Time `gorm:"column:update_time"`
}

func toCompatibilityOsvDO(v *compatibilityOsvDO, data api.Osv, tools, platform []byte) {
	*v = compatibilityOsvDO{
		Architecture:         data.Arch,
		OsVersion:            data.OsVersion,
		OsvName:              data.OsvName,
		Date:                 data.Date,
		OsDownloadLink:       data.OsDownloadLink,
		Type:                 data.Type,
		Details:              data.Details,
		FriendlyLink:         data.FriendlyLink,
		TotalResult:          data.TotalResult,
		CheckSum:             data.CheckSum,
		BaseOpeneulerVersion: data.BaseOpeneulerVersion,
		ToolsResult:          string(tools),
		PlatformResult:       string(platform),
		Updateime:            time.Now(),
	}
}

func (c compatibilityOsvDO) toCompatibilityOsvInfo() (v domain.CompatibilityOsvInfo) {
	v.Id = c.Id
	v.Architecture = c.Architecture
	v.OsVersion = c.OsVersion
	v.OsvName = c.OsvName
	v.Date = c.Date
	v.OsDownloadLink = c.OsDownloadLink
	v.Type = c.Type
	v.Details = c.Details
	v.FriendlyLink = c.FriendlyLink
	v.TotalResult = c.TotalResult
	v.CheckSum = c.CheckSum
	v.BaseOpeneulerVersion = c.BaseOpeneulerVersion
	v.ToolsResult = c.ToolsResult
	v.PlatformResult = c.PlatformResult
	v.Updateime = c.Updateime

	return
}

//
//func (o *OeCompatibilityOsv) GetOsvName() (data []string, err error) {
//	if err = mysqlDb.
//		Model(o).
//		Select("distinct(osv_name) as osvName").
//		Order("osv_name asc").
//		Pluck("osvName", &data).Error; err != nil {
//		return nil, err
//	}
//	return
//}
//
//func (o *OeCompatibilityOsv) GetType() (data []string, err error) {
//	if err = mysqlDb.
//		Model(o).
//		Select("distinct(type) as type").
//		Order("type asc").
//		Pluck("type", &data).Error; err != nil {
//		return nil, err
//	}
//	return
//}
