package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ DataSourceModel = (*customDataSourceModel)(nil)

type (
	// DataSourceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDataSourceModel.
	DataSourceModel interface {
		dataSourceModel
	}

	customDataSourceModel struct {
		*defaultDataSourceModel
	}
)

// NewDataSourceModel returns a model for the database table.
func NewDataSourceModel(conn sqlx.SqlConn) DataSourceModel {
	return &customDataSourceModel{
		defaultDataSourceModel: newDataSourceModel(conn),
	}
}
