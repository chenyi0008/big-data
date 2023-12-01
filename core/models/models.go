package models

import (
	"gorm.io/gorm"
)

type CountryPriceMonth struct {
	gorm.Model
	Day             string  `gorm:"type:varchar(15);"`
	PredictionPrice float64 `gorm:"type:double;"`
	AvgPrice        float64 `gorm:"type:double;"`
}

type CountryProvincePrice struct {
	gorm.Model
	Province string  `gorm:"type:varchar(5);"`
	AvgPrice float64 `gorm:"type:double;"`
}

type DataSource struct {
	//gorm.Model
	//Day         string  `gorm:"type:varchar(15);"`
	//Province    string  `gorm:"type:varchar(5);"`
	//Address     string  `gorm:"type:varchar(15);"`
	//ProductName string  `gorm:"type:varchar(15);"`
	//Category    string  `gorm:"type:varchar(15);"`
	//Price       float64 `gorm:"type:double;"`

	Id          int     `json:"id"`
	Day         string  `json:"day"`
	Province    string  `json:"province"`
	Address     string  `json:"address"`
	ProductName string  `json:"productName"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

func (DataSource) TableName() string {
	return "data_source"
}

type ProvincePriceMonth struct {
	gorm.Model
	Province        string  `gorm:"type:varchar(5);"`
	Day             string  `gorm:"type:varchar(15);"`
	PredictionPrice float64 `gorm:"type:double;"`
	AvgPrice        float64 `gorm:"type:double;"`
}
