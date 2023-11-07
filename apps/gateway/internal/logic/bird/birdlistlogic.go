package bird

import (
	birdClient "birdProtection/apps/bird/client/bird"
	"context"

	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirdListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBirdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirdListLogic {
	return &BirdListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BirdListLogic) BirdList(req *types.BirListReq) (resp *types.BirListResp, err error) {
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
