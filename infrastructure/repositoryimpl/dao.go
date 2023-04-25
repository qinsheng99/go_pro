package repositoryimpl

import (
	"database/sql"

	"gorm.io/gorm"
)

type dbImpl interface {
	InsertTransaction(filter, result interface{}, db *gorm.DB) error
	UpdateTransaction(filter, update interface{}, db *gorm.DB) error
	Begin(...*sql.TxOptions) *gorm.DB
	DeleteTransaction(filter interface{}, db *gorm.DB) error
	Exist(filter interface{}, result interface{}) (bool, error)
}
