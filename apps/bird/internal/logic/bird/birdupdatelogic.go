package birdlogic

import (
	"context"

	"birdProtection/apps/bird/internal/svc"
	"birdProtection/apps/bird/pb/birdservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirdUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBirdUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirdUpdateLogic {
	return &BirdUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BirdUpdateLogic) BirdUpdate(in *birdservice.BirdUpdateReq) (*birdservice.BirdUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &birdservice.BirdUpdateResp{}, nil
}
