package mysql

import (
	"encoding/json"
	"time"

	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/logger"
	"github.com/qinsheng99/go-domain-web/utils"
	"gorm.io/gorm"
)

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

type OsvMapper interface {
	SyncOsv([]api.Osv) error
	OSVFindAll(req api.RequestOsv) (datas []OeCompatibilityOsv, total int64, err error)
}

func NewOsvMapper() OsvMapper {
	return &OeCompatibilityOsv{}
}

func (o *OeCompatibilityOsv) ExistsOsv(version string, tx *gorm.DB) (bool, error) {
	var exists OeCompatibilityOsv
	err := tx.Where("os_version = ?", version).First(&exists).Error
	if err != nil {
		if utils.ErrorNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (o *OeCompatibilityOsv) Delete(version string, tx *gorm.DB) error {
	return tx.Exec("delete from oe_compatibility_osv where os_version = ?", version).Error
}

func (o *OeCompatibilityOsv) UpdateOsv(data *OeCompatibilityOsv, tx *gorm.DB) error {
	return tx.Where("os_version = ?", data.OsVersion).Updates(data).Error
}

func (o *OeCompatibilityOsv) OSVFindAll(req api.RequestOsv) (datas []OeCompatibilityOsv, total int64, err error) {
	q := mysqlDb
	page, size := utils.GetPage(req.Pages)
	query := q.Model(o)
	if req.KeyWord != "" {
		query = query.Where(
			q.Where("osv_name like ?", "%"+req.KeyWord+"%").
				Or("os_version like ?", "%"+req.KeyWord+"%").
				Or("type like ?", "%"+req.KeyWord+"%"),
		)
	}
	if req.OsvName != "" {
		query.Where("osv_name like ?", req.OsvName)
	}

	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}

	if err = query.Count(&total).Error; err != nil {
		logger.Log.Error(err)
		return
	}

	if total == 0 {
		return
	}

	query = query.Order("id desc").Limit(size).Offset((page - 1) * size)
	if err = query.Find(&datas).Error; err != nil {
		logger.Log.Error(err)
		return
	}
	return
}

func (o *OeCompatibilityOsv) GetOsvName() (data []string, err error) {
	if err = mysqlDb.
		Model(o).
		Select("distinct(osv_name) as osvName").
		Order("osv_name asc").
		Pluck("osvName", &data).Error; err != nil {
		return nil, err
	}
	return
}

func (o *OeCompatibilityOsv) GetType() (data []string, err error) {
	if err = mysqlDb.
		Model(o).
		Select("distinct(type) as type").
		Order("type asc").
		Pluck("type", &data).Error; err != nil {
		return nil, err
	}
	return
}

func (o *OeCompatibilityOsv) CreateOsv(data *OeCompatibilityOsv, tx *gorm.DB) error {
	return tx.Create(data).Error
}

func (o *OeCompatibilityOsv) GetOneOSV(osv *OeCompatibilityOsv) (*OeCompatibilityOsv, error) {
	result := mysqlDb.Where(osv).First(osv)
	return osv, result.Error
}

func (o *OeCompatibilityOsv) SyncOsv(osvList []api.Osv) (err error) {
	var tools, platform []byte
	var ok bool
	tx := mysqlDb.Begin()

	for k := range osvList {
		v := osvList[k]
		if len(v.PlatformResult) == 0 && len(v.ToolsResult) == 0 {
			err = o.Delete(v.OsVersion, tx)
			if err != nil {
				tx.Rollback()
				return err
			}
			continue
		}

		tools, err = json.Marshal(v.ToolsResult)
		if err != nil {
			tx.Rollback()
			return err
		}
		platform, err = json.Marshal(v.PlatformResult)
		if err != nil {
			tx.Rollback()
			return err
		}

		data := OeCompatibilityOsv{
			Architecture:         v.Arch,
			OsVersion:            v.OsVersion,
			OsvName:              v.OsvName,
			Date:                 v.Date,
			OsDownloadLink:       v.OsDownloadLink,
			Type:                 v.Type,
			Details:              v.Details,
			FriendlyLink:         v.FriendlyLink,
			TotalResult:          v.TotalResult,
			CheckSum:             v.CheckSum,
			BaseOpeneulerVersion: v.BaseOpeneulerVersion,
			ToolsResult:          string(tools),
			PlatformResult:       string(platform),
			Updateime:            time.Now(),
		}

		if ok, err = o.ExistsOsv(v.OsVersion, tx); err == nil && ok {
			err = o.UpdateOsv(&data, tx)
			if err != nil {
				tx.Rollback()
				return err
			}
		} else if err == nil {
			err = o.CreateOsv(&data, tx)
			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
