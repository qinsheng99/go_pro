package postgresql

import (
	"github.com/qinsheng99/go-domain-web/common/infrastructure/dao"
)

type Config struct {
	dao.DB
	Table table `json:"table"`
}

type table struct {
	CveBasicInfo string `json:"cve_basic_info"`
}
