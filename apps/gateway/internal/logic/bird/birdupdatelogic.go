package bird

import (
	birdClient "birdProtection/apps/bird/client/bird"
	"context"

	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirdUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBirdUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirdUpdateLogic {
	return &BirdUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BirdUpdateLogic) BirdUpdate(req *types.BirUpdateReq) (resp *types.BirUpdateResp, err error) {
	_, err = l.svcCtx.BirdClient.BirdUpdate(l.ctx, &birdClient.BirdUpdateReq{
		BirdID:      req.BirdId,
		BirdName:    req.BirdName,
		BirdType:    req.BirdType,
		Description: req.Description,
		PicUrl:      req.PicUrl,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.BirUpdateResp{
		BirdId: req.BirdId,
	}
	return resp, nil
}
