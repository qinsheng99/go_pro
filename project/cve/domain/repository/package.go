package repository

import "github.com/qinsheng99/go-domain-web/project/cve/domain"

type PkgImpl interface {
	AddApplicationPkg(app []domain.ApplicationPackage) error
}
