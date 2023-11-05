package birdlogic

import (
	"context"

	"birdProtection/apps/bird/internal/svc"
	"birdProtection/apps/bird/pb/birdservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirdListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBirdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirdListLogic {
	return &BirdListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BirdListLogic) BirdList(in *birdservice.BirdListReq) (*birdservice.BirdListResp, error) {
	// todo: add your logic here and delete this line

	return &birdservice.BirdListResp{}, nil
}
