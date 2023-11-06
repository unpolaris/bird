package bird

import (
	"context"

	birdClient "birdProtection/apps/bird/client/bird"

	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBirListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirListLogic {
	return &BirListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BirListLogic) BirList(req *types.BirListReq) (resp *types.BirListResp, err error) {
	res, err := l.svcCtx.BirdClient.BirdList(l.ctx, &birdClient.BirdListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.BirListResp{}
	for _, b := range res.List {
		resp.List = append(resp.List, types.BirListData{
			BirdId:      b.BirdID,
			BirdName:    b.BirdName,
			BirdType:    b.BirdType,
			Description: b.Description,
			PicUrl:      b.PicUrl,
		})
	}
	return resp, nil
}
