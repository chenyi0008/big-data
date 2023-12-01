package logic

import (
	"context"

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

	return
}
