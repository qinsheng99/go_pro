package repositoryimpl

import (
	"gorm.io/gorm"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/dao"
)

type dbimpl interface {
	CreateOrUpdate(tx *gorm.DB, result interface{}, updates ...string) error
	GetRecord(tx *gorm.DB, filter dao.Scope, result interface{}) error
	UpdateRecord(tx *gorm.DB, filter, update interface{}) error

	IsRowNotFound(err error) bool
}
