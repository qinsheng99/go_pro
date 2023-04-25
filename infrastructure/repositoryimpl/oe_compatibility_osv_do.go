package repositoryimpl

import (
	"time"

	"github.com/qinsheng99/go-domain-web/api"
)

type oeCompatibilityOsvDO struct {
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

func toOeCompatibilityOsvDO(v *oeCompatibilityOsvDO, data api.Osv, tools, platform []byte) {
	*v = oeCompatibilityOsvDO{
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

//type OsvMapper interface {
//	OSVFindAll(req domain.OsvDP) (datas []OeCompatibilityOsv, total int64, err error)
//}
//
//func NewOsvMapper() OsvMapper {
//	return &OeCompatibilityOsv{}
//}
//
//func (o *OeCompatibilityOsv) ExistsOsv(version string, tx *gorm.DB) (bool, error) {
//	var exists OeCompatibilityOsv
//	err := tx.Where("os_version = ?", version).First(&exists).Error
//	if err != nil {
//		if utils.ErrorNotFound(err) {
//			return false, nil
//		}
//		return false, err
//	}
//	return true, nil
//}
//
//func (o *OeCompatibilityOsv) Delete(version string, tx *gorm.DB) error {
//	return tx.Exec("delete from oe_compatibility_osv where os_version = ?", version).Error
//}
//
//func (o *OeCompatibilityOsv) UpdateOsv(data *OeCompatibilityOsv, tx *gorm.DB) error {
//	return tx.Where("os_version = ?", data.OsVersion).Updates(data).Error
//}
//
//var mysqlDb = mysql.DB()
//
//func (o *OeCompatibilityOsv) OSVFindAll(req domain.OsvDP) (datas []OeCompatibilityOsv, total int64, err error) {
//	q := mysql.DB()
//	query := q.Model(o)
//	//if req.KeyWord != "" {
//	//	query = query.Where(
//	//		q.Where("osv_name like ?", "%"+req.KeyWord+"%").
//	//			Or("os_version like ?", "%"+req.KeyWord+"%").
//	//			Or("type like ?", "%"+req.KeyWord+"%"),
//	//	)
//	//}
//	//if req.OsvName != "" {
//	//	query.Where("osv_name like ?", req.OsvName)
//	//}
//	//
//	//if req.Type != "" {
//	//	query = query.Where("type = ?", req.Type)
//	//}
//
//	if err = query.Count(&total).Error; err != nil {
//		logger.Log.Error(err)
//		return
//	}
//
//	if total == 0 {
//		return
//	}
//
//	query = query.Order("id desc").Limit(req.Size.Size()).Offset((req.Page.Page() - 1) * req.Size.Size())
//	if err = query.Find(&datas).Error; err != nil {
//		logger.Log.Error(err)
//		return
//	}
//	return
//}
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
//
//func (o *OeCompatibilityOsv) CreateOsv(data *OeCompatibilityOsv, tx *gorm.DB) error {
//	return tx.Create(data).Error
//}
//
//func (o *OeCompatibilityOsv) GetOneOSV(osv *OeCompatibilityOsv) (*OeCompatibilityOsv, error) {
//	result := mysqlDb.Where(osv).First(osv)
//	return osv, result.Error
//}
