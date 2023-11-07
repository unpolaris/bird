package bird

import (
	"birdProtection/pkg/xhttp"
	"net/http"

	"birdProtection/apps/gateway/internal/logic/bird"
	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"
)

func BirdAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BirdAddReq
		if err := xhttp.Parse(r, &req); err != nil {
			xhttp.Error(w, r, err)
			return
		}
		l := bird.NewBirdAddLogic(r.Context(), svcCtx)
		resp, err := l.BirdAdd(&req)
		if err != nil {
			xhttp.Error(w, r, err)
		} else {
			xhttp.OkJson(w, r, resp)
		}
	}
}
