package logic

import (
	"context"

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

	return
}
