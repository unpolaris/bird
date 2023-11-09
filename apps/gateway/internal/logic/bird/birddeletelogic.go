package bird

import (
	birdClient "birdProtection/apps/bird/client/bird"
	"context"

	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirdDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBirdDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirdDeleteLogic {
	return &BirdDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BirdDeleteLogic) BirdDelete(req *types.BirDeleteReq) (resp *types.BirDeleteReq, err error) {
	resp = &types.BirDeleteReq{}
	_, err = l.svcCtx.BirdClient.BirdDelete(l.ctx, &birdClient.BirdDeleteReq{
		BirdID: req.BirdId,
	})
	if err != nil {
		return nil, err
	}
	resp.BirdId = req.BirdId
	return resp, nil
}
