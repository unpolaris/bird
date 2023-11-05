package birdlogic

import (
	"context"

	"birdProtection/apps/bird/internal/svc"
	"birdProtection/apps/bird/pb/birdservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirdAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBirdAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirdAddLogic {
	return &BirdAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BirdAddLogic) BirdAdd(in *birdservice.BirdAddReq) (*birdservice.BirdAddResp, error) {
	// todo: add your logic here and delete this line

	return &birdservice.BirdAddResp{}, nil
}
