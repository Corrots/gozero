package handler

import (
	"net/http"

	"github.com/Corrots/gozero/apps/app/api/internal/logic"
	"github.com/Corrots/gozero/apps/app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HomeBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHomeBannerLogic(r.Context(), svcCtx)
		resp, err := l.HomeBanner(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
