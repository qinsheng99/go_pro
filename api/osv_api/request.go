package osv_api

import "github.com/qinsheng99/go-domain-web/api"

type RequestOsv struct {
	KeyWord string    `json:"keyword"`
	OsvName string    `json:"osvName"`
	Type    string    `json:"type"`
	Pages   api.Pages `json:"pages"`
}

type Osv struct {
	Arch                 string   `json:"arch"`
	OsvName              string   `json:"osv_name"`
	OsVersion            string   `json:"os_version"`
	OsDownloadLink       string   `json:"os_download_link"`
	Type                 string   `json:"type"`
	Date                 string   `json:"date"`
	Details              string   `json:"details"`
	FriendlyLink         string   `json:"friendly_link"`
	TotalResult          string   `json:"total_result"`
	CheckSum             string   `json:"checksum"`
	BaseOpeneulerVersion string   `json:"base_openeuler_version"`
	ToolsResult          []Record `json:"tools_result"`
	PlatformResult       []Record `json:"platform_result"`
}

type Record struct {
	Name    string `json:"name"`
	Percent string `json:"percent"`
	Result  string `json:"result"`
}
