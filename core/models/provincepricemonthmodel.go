package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProvincePriceMonthModel = (*customProvincePriceMonthModel)(nil)

type (
	// ProvincePriceMonthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProvincePriceMonthModel.
	ProvincePriceMonthModel interface {
		provincePriceMonthModel
	}

	customProvincePriceMonthModel struct {
		*defaultProvincePriceMonthModel
	}
)

// NewProvincePriceMonthModel returns a model for the database table.
func NewProvincePriceMonthModel(conn sqlx.SqlConn) ProvincePriceMonthModel {
	return &customProvincePriceMonthModel{
		defaultProvincePriceMonthModel: newProvincePriceMonthModel(conn),
	}
}
