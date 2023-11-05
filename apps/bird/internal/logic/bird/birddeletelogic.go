package birdlogic

import (
	"context"

	"birdProtection/apps/bird/internal/svc"
	"birdProtection/apps/bird/pb/birdservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirdDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBirdDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirdDeleteLogic {
	return &BirdDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BirdDeleteLogic) BirdDelete(in *birdservice.BirdDeleteReq) (*birdservice.BirdDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &birdservice.BirdDeleteResp{}, nil
}
