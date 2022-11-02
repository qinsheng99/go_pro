package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/qinsheng99/go-domain-web/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const CONNMAXLIFTIME = 900

var postgresqlDb *gorm.DB

func Init(cfg *config.PostgresqlConfig) (err error) {
	var sqlDB *sql.DB
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", cfg.DbHost, cfg.DbUser, cfg.DbPwd, cfg.DbName, cfg.DbPort)
	postgresqlDb, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err = postgresqlDb.DB()
	if err != nil {
		return err
	}
	sqlDB.SetConnMaxLifetime(CONNMAXLIFTIME)
	sqlDB.SetMaxOpenConns(cfg.DbMaxConn)
	sqlDB.SetMaxIdleConns(cfg.DbMaxidle)
	return nil
}
