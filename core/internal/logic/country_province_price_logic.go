package logic

import (
	"context"
	"fmt"

	"bigdata/core/internal/svc"
	"bigdata/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountryProvincePriceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountryProvincePriceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountryProvincePriceLogic {
	return &CountryProvincePriceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountryProvincePriceLogic) CountryProvincePrice(req *types.CountryProvincePriceQueryReq) (resp *types.CountryProvincePriceQueryResp, err error) {
	// todo: add your logic here and delete this line
	list := make([]*types.CountryProvincePrice, 0)
	offset := (req.Page - 1) * req.PageSize

	//条件查询
	tx := l.svcCtx.DB.Table("country_province_price").Limit(req.PageSize).Offset(offset)

	if req.Province != "NULL" {
		tx = tx.Where("province = ?", req.Province)
	}
	if req.AvgPrice != 0 {
		tx = tx.Where("ave_price = ?", req.AvgPrice)
	}
	if req.Id != 0 {
		tx = tx.Where("id = ?", req.Id)
	}

	err = tx.Find(&list).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
	resp = &types.CountryProvincePriceQueryResp{
		Data: list,
	}
	return
	return
}
