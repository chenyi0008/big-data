package handler

import (
	"net/http"

	"bigdata/core/internal/logic"
	"bigdata/core/internal/svc"
	"bigdata/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func countryProvincePriceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CountryProvincePriceQueryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCountryProvincePriceLogic(r.Context(), svcCtx)
		resp, err := l.CountryProvincePrice(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
