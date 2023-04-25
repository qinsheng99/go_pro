package mysql

import (
	"database/sql"

	"gorm.io/gorm"

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

func (d MqDao) Insert(filter, result interface{}) error {
	return d.Dao.Insert(filter, result, db)
}

func (d MqDao) InsertTransaction(filter, result interface{}, tx *gorm.DB) error {
	return d.Dao.Insert(filter, result, tx)
}

func (d MqDao) GetRecords(
	filter dao.Scope, result interface{}, p dao.Pagination, sort []dao.SortByColumn,
) error {
	return d.Dao.GetRecords(filter, result, p, sort, db)
}

func (d MqDao) Count(filter dao.Scope) (int, error) {
	return d.Dao.Count(filter, db)
}

func (d MqDao) GetRecord(filter dao.Scope, result interface{}) error {
	return d.Dao.GetRecord(filter, result, db)
}

func (d MqDao) UpdateRecord(filter, update interface{}) error {
	return d.Dao.Update(filter, update, db)
}

func (d MqDao) UpdateTransaction(filter, update interface{}, tx *gorm.DB) error {
	return d.Dao.Update(filter, update, tx)
}

func (d MqDao) Exist(filter interface{}, result interface{}) (bool, error) {
	return d.Dao.Exist(filter, result, db)
}

func (d MqDao) ExecSQL(sql string, result interface{}, args ...interface{}) error {
	return d.Dao.ExecSQL(sql, result, db, args)
}

func (d MqDao) DeleteTransaction(filter interface{}, db *gorm.DB) error {
	return d.Dao.Delete(filter, db)
}

func (d MqDao) Delete(filter interface{}) error {
	return d.Dao.Delete(filter, db)
}
