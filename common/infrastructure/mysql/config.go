package mysql

type Config struct {
	DB
	Table table `json:"table"`
}

type DB struct {
	DbHost    string `json:"db_host"`
	DbPort    int64  `json:"db_port"`
	DbUser    string `json:"db_user"`
	DbPwd     string `json:"db_pwd"`
	DbName    string `json:"db_name"`
	DbMaxConn int    `json:"db_max_conn"`
	DbMaxidle int    `json:"db_maxidle"`
}

type table struct {
	CompatibilityOsv string `json:"compatibility_osv"`
}
