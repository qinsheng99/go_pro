package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

var (
	errRowExists   = errors.New("row exists")
	errRowNotFound = errors.New("row not found")
)

type Scope func(*gorm.DB) *gorm.DB

type SortByColumn struct {
	Column string
	Ascend bool
}

func (s SortByColumn) Order() string {
	v := " ASC"
	if !s.Ascend {
		v = " DESC"
	}
	return s.Column + v
}

type Pagination struct {
	PageNum      int
	CountPerPage int
}

func (p Pagination) Pagination() (limit, offset int) {
	limit = p.CountPerPage

	if limit > 0 && p.PageNum > 0 {
		offset = (p.PageNum - 1) * limit
	}

	return
}

type ColumnFilter struct {
	column string
	symbol string
	value  interface{}
}

func (q *ColumnFilter) condition() string {
	return fmt.Sprintf("%s %s ?", q.column, q.symbol)
}

func NewEqualFilter(column string, value interface{}) ColumnFilter {
	return ColumnFilter{
		column: column,
		symbol: "=",
		value:  value,
	}
}

func NewLikeFilter(column string, value string) ColumnFilter {
	return ColumnFilter{
		column: column,
		symbol: "ilike",
		value:  "%" + value + "%",
	}
}

type Dao struct {
	Name string
}

func (d Dao) Begin(db *gorm.DB, opts ...*sql.TxOptions) *gorm.DB {
	return db.Begin(opts...)
}

func (d Dao) FirstOrCreate(filter, result interface{}, db *gorm.DB) error {
	query := db.Table(d.Name).Where(filter).FirstOrCreate(result)

	if err := query.Error; err != nil {
		return err
	}

	if query.RowsAffected == 0 {
		return errRowExists
	}

	return nil
}

func (d Dao) GetRecords(
	filter Scope, result interface{}, p Pagination, sort []SortByColumn, db *gorm.DB,
) (err error) {
	query := db.Table(d.Name).Scopes(filter)

	var orders []string
	for _, v := range sort {
		orders = append(orders, v.Order())
	}

	if len(orders) >= 0 {
		query.Order(strings.Join(orders, ","))
	}

	if limit, offset := p.Pagination(); limit > 0 {
		query.Limit(limit).Offset(offset)
	}

	err = query.Find(result).Error

	return
}

func (d Dao) Count(filter Scope, db *gorm.DB) (int64, error) {
	var total int64
	query := db.Table(d.Name).Scopes(filter)

	err := query.Count(&total).Error

	return total, err
}

func (d Dao) GetRecord(filter Scope, result interface{}, db *gorm.DB) error {
	err := db.Table(d.Name).Scopes(filter).First(result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errRowNotFound
	}

	return err
}

func (d Dao) Update(filter, update interface{}, db *gorm.DB) (err error) {
	query := db.Table(d.Name).Where(filter).Updates(update)
	if err = query.Error; err != nil {
		return
	}

	if query.RowsAffected == 0 {
		err = errRowNotFound
	}

	return
}

func (d Dao) Exist(filter interface{}, result interface{}, db *gorm.DB) (bool, error) {
	err := db.Table(d.Name).Where(filter).First(result).Error
	if err == nil {
		return true, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return false, err
}

func (d Dao) Delete(filter interface{}, db *gorm.DB) error {
	return db.Table(d.Name).Where(filter).Delete(nil).Error
}

func (d Dao) ExecSQL(sql string, result interface{}, db *gorm.DB, args ...interface{}) error {
	return db.Exec(sql, args...).Find(result).Error
}

func (d Dao) IsRowNotFound(err error) bool {
	return errors.Is(err, errRowNotFound)
}

func (d Dao) IsRowExists(err error) bool {
	return errors.Is(err, errRowExists)
}
