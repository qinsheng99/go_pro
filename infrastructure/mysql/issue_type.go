package mysql

type IssueType struct {
	Id           int64  `gorm:"column:id" json:"id"`
	UniqueId     int64  `gorm:"column:unique_id" json:"uniqueId"`
	Name         string `gorm:"column:name" json:"name"`
	Platform     string `gorm:"column:platform" json:"platform"`
	Organization string `gorm:"column:organization" json:"organization"`
	Template     string `gorm:"column:template" json:"template"`
}

func (i *IssueType) TableName() string {
	return "issue_type"
}

func (i *IssueType) List() (data []IssueType, err error) {
	err = Getmysqldb().Find(&data).Error
	return
}

func (i *IssueType) Find() (err error) {
	err = Getmysqldb().Where("unique_id = ?", i.UniqueId).First(i).Error
	return
}
