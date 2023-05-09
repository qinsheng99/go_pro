package domain

import (
	"time"

	"github.com/qinsheng99/go-domain-web/project/sort/domain/dp"
)

type OsvOptions struct {
	KeyWord string
	OsvName string
	Type    string

	Page dp.Page
	Size dp.Size
}

type CompatibilityOsvInfo struct {
	Id                   int64     `json:"id"`
	Architecture         string    `json:"arch"`
	OsVersion            string    `json:"osVersion"`
	OsvName              string    `json:"osvName"`
	Date                 string    `json:"date"`
	OsDownloadLink       string    `json:"osDownloadLink"`
	Type                 string    `json:"type"`
	Details              string    `json:"details"`
	FriendlyLink         string    `json:"friendlyLink"`
	TotalResult          string    `json:"totalResult"`
	CheckSum             string    `json:"checksum"`
	BaseOpeneulerVersion string    `json:"baseOpeneulerVersion"`
	ToolsResult          string    `json:"toolsResult"`
	PlatformResult       string    `json:"platformResult"`
	Updateime            time.Time `json:"updateTime"`
}
