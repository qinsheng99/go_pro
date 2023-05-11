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

func (d PgDao) Insert(result interface{}) error {
	return db.Table(d.Dao.Name).Create(result).Error
}

func (d PgDao) FirstOrCreate(filter, result interface{}) error {
	return d.Dao.FirstOrCreate(filter, result, db)
}

// CreateOrUpdate use uuid if you find record can update
func (d PgDao) CreateOrUpdate(result interface{}, updates ...string) error {
	return db.Table(d.Dao.Name).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uuid"}},
		DoUpdates: clause.AssignmentColumns(updates),
	}).Create(result).Error
}

func (d PgDao) InsertTransaction(filter, result interface{}, tx *gorm.DB) error {
	return d.Dao.FirstOrCreate(filter, result, tx)
}

func (d PgDao) Transaction(f func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return db.Transaction(f, opts...)
}

func (d PgDao) GetRecords(
	filter dao.Scope, result interface{}, p dao.Pagination, sort []dao.SortByColumn,
) error {
	return d.Dao.GetRecords(filter, result, p, sort, db)
}

func (d PgDao) Count(filter dao.Scope) (int64, error) {
	return d.Dao.Count(filter, db)
}

func (d PgDao) GetRecord(filter dao.Scope, result interface{}) error {
	return d.Dao.GetRecord(filter, result, db)
}

func (d PgDao) UpdateRecord(filter, update interface{}) error {
	return d.Dao.Update(filter, update, db)
}

func (d PgDao) UpdateTransaction(filter, update interface{}, tx *gorm.DB) error {
	return d.Dao.Update(filter, update, tx)
}

func (d PgDao) Exist(filter interface{}, result interface{}) (bool, error) {
	return d.Dao.Exist(filter, result, db)
}

func (d PgDao) ExecSQL(sql string, result interface{}, args ...interface{}) error {
	return d.Dao.ExecSQL(sql, result, db, args)
}

func (d PgDao) DeleteTransaction(filter interface{}, tx *gorm.DB) error {
	return d.Dao.Delete(filter, tx)
}

func (d PgDao) Delete(filter interface{}) error {
	return d.Dao.Delete(filter, db)
}
