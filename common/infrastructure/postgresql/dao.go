package postgresql

import (
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/dao"
)

type PgDao struct {
	dao.Dao
}

func NewPgDao(name string) PgDao {
	return PgDao{Dao: dao.Dao{Name: name}}
}

func (d PgDao) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return d.Dao.Begin(db, opts...)
}

func (d PgDao) Insert(tx *gorm.DB, result interface{}) error {
	return d.checkDB(tx).Table(d.Dao.Name).Create(result).Error
}

func (d PgDao) FirstOrCreate(tx *gorm.DB, filter, result interface{}) error {
	return d.Dao.FirstOrCreate(filter, result, d.checkDB(tx))
}

func (d PgDao) CreateOrUpdate(tx *gorm.DB, result interface{}, updates ...string) error {
	return d.checkDB(tx).Table(d.Dao.Name).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uuid"}},
		DoUpdates: clause.AssignmentColumns(updates),
	}).Create(result).Error
}

func (d PgDao) GetRecords(
	tx *gorm.DB, filter dao.Scope, result interface{}, p dao.Pagination, sort []dao.SortByColumn,
) error {
	return d.Dao.GetRecords(filter, result, p, sort, d.checkDB(tx))
}

func (d PgDao) Count(tx *gorm.DB, filter dao.Scope) (int64, error) {
	return d.Dao.Count(filter, d.checkDB(tx))
}

func (d PgDao) GetRecord(tx *gorm.DB, filter dao.Scope, result interface{}) error {
	return d.Dao.GetRecord(filter, result, d.checkDB(tx))
}

func (d PgDao) Transaction(tx *gorm.DB, f func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return d.checkDB(tx).Transaction(f, opts...)
}

func (d PgDao) UpdateRecord(tx *gorm.DB, filter, update interface{}) error {
	return d.Dao.Update(filter, update, d.checkDB(tx))
}

func (d PgDao) Exist(tx *gorm.DB, filter interface{}, result interface{}) (bool, error) {
	return d.Dao.Exist(filter, result, d.checkDB(tx))
}

func (d PgDao) ExecSQL(tx *gorm.DB, sql string, result interface{}, args ...interface{}) error {
	return d.Dao.ExecSQL(sql, result, d.checkDB(tx), args)
}

func (d PgDao) Delete(filter interface{}, tx *gorm.DB) error {
	return d.Dao.Delete(filter, d.checkDB(tx))
}

func (d PgDao) DB() *gorm.DB {
	return db
}

func (d PgDao) checkDB(gd *gorm.DB) *gorm.DB {
	if gd != nil {
		return gd
	}

	return db
}
