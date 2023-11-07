package bird

import (
	"birdProtection/pkg/xhttp"
	"net/http"

	"birdProtection/apps/gateway/internal/logic/bird"
	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"
)

// BirdListHandler TODO: Comments
// @Tags TODO: Tags
// @Summary BirdListHandler
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body types.BirListReq true TODO: "Comments"
// @Success 200 {object} types.Response{data=} "{"code":200,"msg":"OK","data":{}}"
// @Router
func BirdListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BirListReq
		if err := xhttp.Parse(r, &req); err != nil {
			xhttp.Error(w, r, err)
			return
		}

		l := bird.NewBirdListLogic(r.Context(), svcCtx)
		resp, err := l.BirdList(&req)
		if err != nil {
			xhttp.Error(w, r, err)
		} else {
			xhttp.OkJson(w, r, resp)
		}
	}
}
