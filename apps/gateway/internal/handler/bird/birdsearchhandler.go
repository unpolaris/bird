package bird

import (
	"birdProtection/pkg/xhttp"
	"net/http"

	"birdProtection/apps/gateway/internal/logic/bird"
	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"
)

// BirdSearchHandler TODO: Comments
// @Tags TODO: Tags
// @Summary BirdSearchHandler
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body types.BirdSearchReq true TODO: "Comments"
// @Success 200 {object} types.Response{data=} "{"code":200,"msg":"OK","data":{}}"
// @Router
func BirdSearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BirdSearchReq
		if err := xhttp.Parse(r, &req); err != nil {
			xhttp.Error(w, r, err)
			return
		}

		l := bird.NewBirdSearchLogic(r.Context(), svcCtx)
		resp, err := l.BirdSearch(&req)
		if err != nil {
			xhttp.Error(w, r, err)
		} else {
			xhttp.OkJson(w, r, resp)
		}
	}
}
