package bird

import (
	"net/http"

	"birdProtection/apps/gateway/internal/logic/bird"
	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BirGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BirGetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := bird.NewBirGetLogic(r.Context(), svcCtx)
		resp, err := l.BirGet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
