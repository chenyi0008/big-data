package logic

import (
	"context"

	"bigdata/core/internal/svc"
	"bigdata/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountryPriceMonthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountryPriceMonthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountryPriceMonthLogic {
	return &CountryPriceMonthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountryPriceMonthLogic) CountryPriceMonth(req *types.CountryPriceMonthQueryReq) (resp *types.CountryPriceMonthQueryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
