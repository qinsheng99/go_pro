package repositoryimpl

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/dao"
)

type dbimpl interface {
	Insert(tx *gorm.DB, result interface{}) error
	FirstOrCreate(tx *gorm.DB, filter, result interface{}) error
	CreateOrUpdate(tx *gorm.DB, result interface{}, updates ...string) error
	GetRecord(tx *gorm.DB, filter dao.Scope, result interface{}) error
	GetRecords(
		tx *gorm.DB, filter dao.Scope, result interface{}, p dao.Pagination, sort []dao.SortByColumn,
	) error
	UpdateRecord(tx *gorm.DB, filter, update interface{}) error
	Delete(tx *gorm.DB, filter interface{}) error

	Transaction(tx *gorm.DB, f func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
	DB() *gorm.DB

	IsRowNotFound(err error) bool
}
