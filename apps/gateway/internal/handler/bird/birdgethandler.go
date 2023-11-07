package bird

import (
	"birdProtection/apps/gateway/internal/logic/bird"
	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"
	"birdProtection/pkg/xhttp"
	"net/http"
)

// BirdGetHandler TODO: Comments
// @Tags TODO: Tags
// @Summary BirdGetHandler
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body types.BirGetReq true TODO: "Comments"
// @Success 200 {object} types.Response{data=} "{"code":200,"msg":"OK","data":{}}"
// @Router
func BirdGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BirGetReq
		if err := xhttp.Parse(r, &req); err != nil {
			xhttp.Error(w, r, err)
			return
		}

		l := bird.NewBirdGetLogic(r.Context(), svcCtx)
		resp, err := l.BirdGet(&req)
		if err != nil {
			xhttp.Error(w, r, err)
		} else {
			xhttp.OkJson(w, r, resp)
		}
	}
}
