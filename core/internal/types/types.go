// Code generated by goctl. DO NOT EDIT.
package types

type TestReq struct {
	Id int `form:"id"`
}

type TestResp struct {
	Msg       string `json:"msg"`
	Data      string `json:"data"`
	IsSeccess bool   `json:"isSeccess"`
}

type CountryPriceMonthQueryReq struct {
	Page            int     `json:"page"`
	PageSize        int     `json:"pageSize"`
	Id              int     `json:"id,default=0"`
	Day             string  `json:"day,default=NULL"`
	PredictionPrice float64 `json:"predictionPrice,default=0"`
	AvgPrice        float64 `json:"avgPrice,default=0"`
}

type CountryPriceMonthQueryResp struct {
	IsSuccess bool                 `json:"isSuccess"`
	Data      []*CountryPriceMonth `json:"data"`
}

type CountryPriceMonth struct {
	Id               int     `json:"id"`
	Day              string  `json:"day"`
	Prediction_price float64 `json:"predictionPrice"`
	AvgPrice         float64 `json:"avgPrice"`
}

type CountryProvincePriceQueryReq struct {
	Page     int     `json:"page"`
	PageSize int     `json:"pageSize"`
	Id       int     `json:"id,default=0"`
	Province string  `json:"province,default=NULL"`
	AvgPrice float64 `json:"avgPrice,default=0"`
}

type CountryProvincePriceQueryResp struct {
	IsSuccess bool                    `json:"isSuccess"`
	Data      []*CountryProvincePrice `json:"data"`
}

type CountryProvincePrice struct {
	Id       int     `json:"id"`
	Province string  `json:"province"`
	AvgPrice float64 `json:"avgPrice"`
}

type DataSourceQueryReq struct {
	Page        int     `json:"page"`
	PageSize    int     `json:"pageSize"`
	Id          int     `json:"id,default=0"`
	Day         string  `json:"day,default=NULL"`
	Province    string  `json:"province,default=NULL"`
	Address     string  `json:"address,default=NULL"`
	ProductName string  `json:"productName,default=NULL"`
	Category    string  `json:"category,default=NULL"`
	Price       float64 `json:"price,default=0.0"`
}

type DataSourceQueryResp struct {
	IsSuccess bool          `json:"isSuccess"`
	Data      []*DataSource `json:"data"`
}

type DataSource struct {
	Id          int     `json:"id"`
	Day         string  `json:"day"`
	Province    string  `json:"province"`
	Address     string  `json:"address"`
	ProductName string  `json:"productName"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

type ProvincePriceMonthQueryReq struct {
	Page            int     `json:"page"`
	PageSize        int     `json:"pageSize"`
	Id              int     `json:"id,default=0"`
	Province        string  `json:"province,default=NULL"`
	Day             string  `json:"day,default=NULL"`
	PredictionPrice float64 `json:"predictionPrice,default=0"`
	AvgPrice        float64 `json:"avgPrice,default=0"`
}

type ProvincePriceMonthQueryResp struct {
	IsSuccess bool                  `json:"isSuccess"`
	Data      []*ProvincePriceMonth `json:"data"`
}

type ProvincePriceMonth struct {
	Id              int     `json:"id"`
	Province        string  `json:"province"`
	Day             string  `json:"day"`
	PredictionPrice float64 `json:"predictionPrice"`
	AvgPrice        float64 `json:"avgPrice"`
}

type AxisQueryReq struct {
	Type string `form:"type"`
}

type AxisQueryResp struct {
	IsSuccess bool      `json:"isSuccess"`
	Data      []*string `json:"data"`
}
