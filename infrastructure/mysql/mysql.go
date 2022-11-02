package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qinsheng99/go-domain-web/config"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDb *gorm.DB

const CONNMAXLIFTIME = 900

func Init(cfg *config.MysqlConfig) (err error) {
	var sqlDB *sql.DB
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPwd, cfg.DbHost, cfg.DbPort, cfg.DbName)
	mysqlDb, err = gorm.Open(gormmysql.New(gormmysql.Config{
		DSN:                       dsn,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err = mysqlDb.DB()
	if err != nil {
		return err
	}
	sqlDB.SetConnMaxLifetime(CONNMAXLIFTIME)
	sqlDB.SetMaxOpenConns(cfg.DbMaxConn)
	sqlDB.SetMaxIdleConns(cfg.DbMaxidle)
	return nil
}
