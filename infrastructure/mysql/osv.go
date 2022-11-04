package mysql

import "time"

type OeCompatibilityOsv struct {
	Id                   int64     `gorm:"column:id" json:"id"`
	Architecture         string    `gorm:"column:architecture" json:"arch"`
	OsVersion            string    `gorm:"column:os_version" json:"osVersion"`
	OsvName              string    `gorm:"column:osv_name" json:"osvName"`
	Date                 string    `gorm:"column:date" json:"date"`
	OsDownloadLink       string    `gorm:"column:os_download_link" json:"osDownloadLink"`
	Type                 string    `gorm:"column:type" json:"type"`
	Details              string    `gorm:"column:details" json:"details"`
	FriendlyLink         string    `gorm:"column:friendly_link" json:"friendlyLink"`
	TotalResult          string    `gorm:"column:total_result" json:"totalResult"`
	CheckSum             string    `gorm:"column:checksum" json:"checksum"`
	BaseOpeneulerVersion string    `gorm:"column:base_openeuler_version" json:"baseOpeneulerVersion"`
	ToolsResult          string    `gorm:"column:tools_result" json:"toolsResult"`
	PlatformResult       string    `gorm:"column:platform_result" json:"platformResult"`
	Updateime            time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (o *OeCompatibilityOsv) TableName() string {
	return "oe_compatibility_osv"
}

type ROeCompatibilityOsv struct {
	OeCompatibilityOsv
	Updateime      string   `json:"updateTime"`
	ToolsResult    []Record `gorm:"column:tools_result" json:"toolsResult"`
	PlatformResult []Record `gorm:"column:platform_result" json:"platformResult"`
}
type Record struct {
	Name    string `json:"name"`
	Percent string `json:"percent"`
	Result  string `json:"result"`
}
