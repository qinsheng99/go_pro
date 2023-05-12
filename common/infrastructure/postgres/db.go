package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const CONNMAXLIFTIME = 900

var db *gorm.DB

func Init(cfg *Config) (err error) {
	var sqlDB *sql.DB
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", cfg.DbHost, cfg.DbUser, cfg.DbPwd, cfg.DbName, cfg.DbPort)
	l := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{Logger: l})
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
