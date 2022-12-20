package mysql

import (
	"time"
)

type Email struct {
	Id         int64     `gorm:"column:id" json:"id"`
	Email      string    `gorm:"column:email" json:"email"`
	Code       string    `gorm:"column:code" json:"code"`
	IsDelete   int       `gorm:"column:is_delete" json:"-"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
}

func (e *Email) TableName() string {
	return "email"
}

func (e *Email) Check() bool {
	err := Getmysqldb().Where("email = ? and code = ? and is_delete = 0", e.Email, e.Code).First(e).Error
	if err != nil {
		return false
	}

	return true
}

func (e *Email) DeleteCode() {
	Getmysqldb().Model(e).Where("id = ?", e.Id).Update("is_delete", 1)
}

func (e *Email) Insert() error {
	return Getmysqldb().Create(e).Error
}

func (e *Email) Update() error {
	return Getmysqldb().Model(e).Where("id = ?", e.Id).UpdateColumns(map[string]interface {
	}{"code": e.Code, "create_time": e.CreateTime, "is_delete": 0}).Error
}

func (e *Email) Exist() bool {
	err := Getmysqldb().Where("email = ?", e.Email).First(e).Error
	if err != nil {
		return false
	}

	return true
}
