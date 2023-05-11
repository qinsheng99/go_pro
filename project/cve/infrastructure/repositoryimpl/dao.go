package repositoryimpl

import (
	"github.com/qinsheng99/go-domain-web/common/infrastructure/dao"
)

type dbimpl interface {
	CreateOrUpdate(result interface{}, updates ...string) error
	GetRecord(filter dao.Scope, result interface{}) error
	UpdateRecord(filter, update interface{}) error

	IsRowNotFound(err error) bool
}
