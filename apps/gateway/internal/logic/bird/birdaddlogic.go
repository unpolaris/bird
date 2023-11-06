package bird

import (
	birdClient "birdProtection/apps/bird/client/bird"
	"context"

	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirdAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBirdAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirdAddLogic {
	return &BirdAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BirdAddLogic) BirdAdd(req *types.BirdAddReq) (resp *types.BirdAddResp, err error) {
	_, err = l.svcCtx.BirdClient.BirdAdd(l.ctx, &birdClient.BirdAddReq{
		BirdName:    req.BirdName,
		BirdType:    req.BirdType,
		Description: req.Description,
		PicUrl:      req.PicUrl,
	})
	if err != nil {
		return nil, err
	}
	resp.BirdName = req.BirdName
	return resp, nil
}
