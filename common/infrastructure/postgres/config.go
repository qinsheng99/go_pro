package postgres

import (
	"github.com/qinsheng99/go-domain-web/common/infrastructure/dao"
)

type Config struct {
	dao.DB
	Table table `json:"table"`
}

type table struct {
	BasicInfo      string `json:"basic_info"`
	BasePkg        string `json:"base_pkg"`
	ApplicationPkg string `json:"application_pkg"`
}
