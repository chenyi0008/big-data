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
	//条件查询

	tx := l.svcCtx.DB.Table("data_source").Limit(req.PageSize).Offset(offset)
	if req.Id != 0 {
		tx = tx.Where("id = ?", req.Id)
	}
	if req.Address != "NULL" {
		tx = tx.Where("address = ?", req.Address)
	}
	if req.Price != 0 {
		tx = tx.Where("price = ?", req.Price)
	}
	if req.Category != "NULL" {
		tx = tx.Where("category = ?", req.Category)
	}
	if req.Province != "NULL" {
		tx = tx.Where("province = ?", req.Province)
	}
	if req.ProductName != "NULL" {
		tx = tx.Where("product_name = ?", req.ProductName)
	}
	err = tx.Find(&list).Error

	if err != nil {
		panic(err)
	}
	fmt.Println(list)
	resp = &types.DataSourceQueryResp{
		Data: list,
	}
	return
}
