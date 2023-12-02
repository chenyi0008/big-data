package logic

import (
	"context"
	"fmt"

	"bigdata/core/internal/svc"
	"bigdata/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProvincePriceMonthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProvincePriceMonthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProvincePriceMonthLogic {
	return &ProvincePriceMonthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProvincePriceMonthLogic) ProvincePriceMonth(req *types.ProvincePriceMonthQueryReq) (resp *types.ProvincePriceMonthQueryResp, err error) {
	// todo: add your logic here and delete this line
	list := make([]*types.ProvincePriceMonth, 0)
	offset := (req.Page - 1) * req.PageSize

	//条件查询
	tx := l.svcCtx.DB.Table("province_price_month").Limit(req.PageSize).Offset(offset)

	if req.Id != 0 {
		tx = tx.Where("id = ?", req.Id)
	}
	if req.Province != "NULL" {
		tx = tx.Where("province = ?", req.Province)
	}
	if req.Day != "NULL" {
		tx = tx.Where("day = ?", req.Day)
	}
	if req.PredictionPrice != 0 {
		tx = tx.Where("prediction_price = ?", req.PredictionPrice)
	}
	if req.AvgPrice != 0 {
		tx = tx.Where("avg_price = ?", req.AvgPrice)
	}

	err = tx.Find(&list).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
	resp = &types.ProvincePriceMonthQueryResp{
		Data:      list,
		IsSuccess: true,
	}
	return

}
