package birdlogic

import (
	"context"

	"birdProtection/apps/bird/internal/svc"
	"birdProtection/apps/bird/pb/birdservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirdInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBirdInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirdInfoLogic {
	return &BirdInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BirdInfoLogic) BirdInfo(in *birdservice.BirdInfoReq) (*birdservice.BirdInfoResp, error) {
	// todo: add your logic here and delete this line

	return &birdservice.BirdInfoResp{}, nil
}
