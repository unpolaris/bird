package bird

import (
	birdClient "birdProtection/apps/bird/client/bird"
	"context"

	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirdGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBirdGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirdGetLogic {
	return &BirdGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BirdGetLogic) BirdGet(req *types.BirGetReq) (resp *types.BirGetResp, err error) {
	res, err := l.svcCtx.BirdClient.BirdInfo(l.ctx, &birdClient.BirdInfoReq{
		BirdID: req.BirdId,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.BirGetResp{
		BirdId:      res.BirdID,
		BirdName:    res.BirdName,
		BirdType:    res.BirdType,
		Description: res.Description,
		PicUrl:      res.PicUrl,
		CreateTime:  res.CreateTime,
		UpdateTime:  res.UpdateTime,
	}
	return resp, nil
}
