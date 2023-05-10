package postgresql

import "github.com/qinsheng99/go-domain-web/common/infrastructure/mysql"

type Config struct {
	mysql.DB
	Table table `json:"table"`
}

type table struct {
	CveBasicInfoTest string `json:"cve_basic_info_test"`
}
