package logic

import (
	"context"
	"fmt"

	"bigdata/core/internal/svc"
	"bigdata/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DataSourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDataSourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DataSourceLogic {
	return &DataSourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DataSourceLogic) DataSource(req *types.DataSourceQueryReq) (resp *types.DataSourceQueryResp, err error) {
	// todo: add your logic here and delete this line
	list := make([]*types.DataSource, 0)
	offset := (req.Page - 1) * req.PageSize
	err = l.svcCtx.DB.Table("data_source").Limit(req.PageSize).Offset(offset).Find(&list).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
	resp = &types.DataSourceQueryResp{
		Data: list,
	}
	return
}
