package handler

import (
	"net/http"

	"bigdata/core/internal/logic"
	"bigdata/core/internal/svc"
	"bigdata/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func dataSourceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DataSourceQueryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDataSourceLogic(r.Context(), svcCtx)
		resp, err := l.DataSource(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
