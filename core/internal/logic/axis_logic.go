package logic

import (
	"context"

	"bigdata/core/internal/svc"
	"bigdata/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AxisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAxisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AxisLogic {
	return &AxisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AxisLogic) Axis(req *types.AxisQueryReq) (resp *types.AxisQueryResp, err error) {

	list := make([]*string, 0)

	err = l.svcCtx.DB.Table("data_source").Distinct(req.Type).Pluck(req.Type, &list).Error
	if err != nil {
		panic(err)
	}

	resp = &types.AxisQueryResp{
		Data:      list,
		IsSuccess: true,
	}
	return
}
