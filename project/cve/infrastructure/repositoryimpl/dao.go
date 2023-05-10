package repositoryimpl

import (
	"github.com/qinsheng99/go-domain-web/common/infrastructure/dao"
)

type dbimpl interface {
	FirstOrCreate(filter, result interface{}) error
	GetRecord(filter dao.Scope, result interface{}) error
	UpdateRecord(filter, update interface{}) error

	IsRowNotFound(err error) bool
}
