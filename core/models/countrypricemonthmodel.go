package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CountryPriceMonthModel = (*customCountryPriceMonthModel)(nil)

type (
	// CountryPriceMonthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCountryPriceMonthModel.
	CountryPriceMonthModel interface {
		countryPriceMonthModel
	}

	customCountryPriceMonthModel struct {
		*defaultCountryPriceMonthModel
	}
)

// NewCountryPriceMonthModel returns a model for the database table.
func NewCountryPriceMonthModel(conn sqlx.SqlConn) CountryPriceMonthModel {
	return &customCountryPriceMonthModel{
		defaultCountryPriceMonthModel: newCountryPriceMonthModel(conn),
	}
}
