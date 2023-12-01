// Code generated by goctl. DO NOT EDIT.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	countryProvincePriceFieldNames          = builder.RawFieldNames(&CountryProvincePrice{})
	countryProvincePriceRows                = strings.Join(countryProvincePriceFieldNames, ",")
	countryProvincePriceRowsExpectAutoSet   = strings.Join(stringx.Remove(countryProvincePriceFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	countryProvincePriceRowsWithPlaceHolder = strings.Join(stringx.Remove(countryProvincePriceFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	countryProvincePriceModel interface {
		Insert(ctx context.Context, data *CountryProvincePrice) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*CountryProvincePrice, error)
		Update(ctx context.Context, data *CountryProvincePrice) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCountryProvincePriceModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CountryProvincePrice struct {
		Id       int64           `db:"id"`
		Province sql.NullString  `db:"province"`
		AvgPrice sql.NullFloat64 `db:"avg_price"`
	}
)

func newCountryProvincePriceModel(conn sqlx.SqlConn) *defaultCountryProvincePriceModel {
	return &defaultCountryProvincePriceModel{
		conn:  conn,
		table: "`country_province_price`",
	}
}

func (m *defaultCountryProvincePriceModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCountryProvincePriceModel) FindOne(ctx context.Context, id int64) (*CountryProvincePrice, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", countryProvincePriceRows, m.table)
	var resp CountryProvincePrice
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCountryProvincePriceModel) Insert(ctx context.Context, data *CountryProvincePrice) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, countryProvincePriceRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Province, data.AvgPrice)
	return ret, err
}

func (m *defaultCountryProvincePriceModel) Update(ctx context.Context, data *CountryProvincePrice) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, countryProvincePriceRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Province, data.AvgPrice, data.Id)
	return err
}

func (m *defaultCountryProvincePriceModel) tableName() string {
	return m.table
}