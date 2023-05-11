package repositoryimpl

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/dao"
)

type dbImpl interface {
	Insert(result interface{}, tx *gorm.DB) error
	UpdateRecord(tx *gorm.DB, filter, update interface{}) error
	FirstOrCreate(tx *gorm.DB, filter, result interface{}) error
	CreateOrUpdate(result interface{}, tx *gorm.DB, updates ...string) error

	Transaction(tx *gorm.DB, f func(tx *gorm.DB) error, opts ...*sql.TxOptions) error

	GetRecords(tx *gorm.DB, filter dao.Scope, result interface{}, p dao.Pagination, sort []dao.SortByColumn) error
	Count(tx *gorm.DB, filter dao.Scope) (int64, error)

	Begin(...*sql.TxOptions) *gorm.DB
	Delete(filter interface{}, tx *gorm.DB) error
	Exist(tx *gorm.DB, filter interface{}, result interface{}) (bool, error)
}
