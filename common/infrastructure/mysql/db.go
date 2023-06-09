package mysql

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm/logger"

	_ "github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

const CONNMAXLIFTIME = 900

func Init(cfg *Config) (err error) {
	var sqlDB *sql.DB
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPwd, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err = gorm.Open(gormmysql.New(gormmysql.Config{
		DSN:                       dsn,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return err
	}
	sqlDB, err = db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetConnMaxLifetime(CONNMAXLIFTIME)
	sqlDB.SetMaxOpenConns(cfg.DbMaxConn)
	sqlDB.SetMaxIdleConns(cfg.DbMaxidle)
	return nil
}
