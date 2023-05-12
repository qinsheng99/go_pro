package mysql

import (
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/dao"
)

type MqDao struct {
	dao.Dao
}

func NewMqDao(name string) MqDao {
	return MqDao{Dao: dao.Dao{Name: name}}
}

func (d MqDao) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return d.Dao.Begin(db, opts...)
}

func (d MqDao) Insert(tx *gorm.DB, result interface{}) error {
	return d.checkDB(tx).Table(d.Dao.Name).Create(result).Error
}

func (d MqDao) FirstOrCreate(tx *gorm.DB, filter, result interface{}) error {
	return d.Dao.FirstOrCreate(filter, result, d.checkDB(tx))
}

func (d MqDao) CreateOrUpdate(tx *gorm.DB, result interface{}, updates ...string) error {
	return d.checkDB(tx).Table(d.Dao.Name).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns(updates),
	}).Create(result).Error
}

func (d MqDao) GetRecords(
	tx *gorm.DB, filter dao.Scope, result interface{}, p dao.Pagination, sort []dao.SortByColumn,
) error {
	return d.Dao.GetRecords(filter, result, p, sort, d.checkDB(tx))
}

func (d MqDao) Count(tx *gorm.DB, filter dao.Scope) (int64, error) {
	return d.Dao.Count(filter, d.checkDB(tx))
}

func (d MqDao) GetRecord(tx *gorm.DB, filter dao.Scope, result interface{}) error {
	return d.Dao.GetRecord(filter, result, d.checkDB(tx))
}

func (d MqDao) Transaction(tx *gorm.DB, f func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return d.checkDB(tx).Transaction(f, opts...)
}

func (d MqDao) UpdateRecord(tx *gorm.DB, filter, update interface{}) error {
	return d.Dao.Update(filter, update, d.checkDB(tx))
}

func (d MqDao) Exist(tx *gorm.DB, filter interface{}, result interface{}) (bool, error) {
	return d.Dao.Exist(filter, result, d.checkDB(tx))
}

func (d MqDao) ExecSQL(tx *gorm.DB, sql string, result interface{}, args ...interface{}) error {
	return d.Dao.ExecSQL(sql, result, d.checkDB(tx), args)
}

func (d MqDao) Delete(tx *gorm.DB, filter interface{}) error {
	return d.Dao.Delete(filter, d.checkDB(tx))
}

func (d MqDao) DB() *gorm.DB {
	return db
}

func (d MqDao) checkDB(gd *gorm.DB) *gorm.DB {
	if gd != nil {
		return gd
	}

	return db
}
