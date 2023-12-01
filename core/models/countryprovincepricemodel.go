package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CountryProvincePriceModel = (*customCountryProvincePriceModel)(nil)

type (
	// CountryProvincePriceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCountryProvincePriceModel.
	CountryProvincePriceModel interface {
		countryProvincePriceModel
	}

	customCountryProvincePriceModel struct {
		*defaultCountryProvincePriceModel
	}
)

// NewCountryProvincePriceModel returns a model for the database table.
func NewCountryProvincePriceModel(conn sqlx.SqlConn) CountryProvincePriceModel {
	return &customCountryProvincePriceModel{
		defaultCountryProvincePriceModel: newCountryProvincePriceModel(conn),
	}
}
