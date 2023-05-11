package mysql

import "github.com/qinsheng99/go-domain-web/common/infrastructure/dao"

type Config struct {
	dao.DB
	Table table `json:"table"`
}

type table struct {
	CompatibilityOsv string `json:"compatibility_osv"`
}
