package repositoryimpl

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/dao"
)

type dbImpl interface {
	InsertTransaction(filter, result interface{}, db *gorm.DB) error
	UpdateTransaction(filter, update interface{}, db *gorm.DB) error

	GetRecords(dao.Scope, interface{}, dao.Pagination, []dao.SortByColumn) error
	Count(dao.Scope) (int64, error)

	Begin(...*sql.TxOptions) *gorm.DB
	DeleteTransaction(filter interface{}, db *gorm.DB) error
	Exist(filter interface{}, result interface{}) (bool, error)
}
