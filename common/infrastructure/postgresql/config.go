package postgresql

import "github.com/qinsheng99/go-domain-web/common/infrastructure/mysql"

type Config struct {
	mysql.DB
	Table table `json:"table"`
}

type table struct {
	CveOriginRecord string `json:"cve_origin_record"`
}
