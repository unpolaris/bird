package birdlogic

import (
	"birdProtection/model"
	"birdProtection/pkg/errcode"
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
	var (
		birdDB = model.NewBirdModel(l.svcCtx.BirdDB, l.ctx)
		resp   = &birdservice.BirdAddResp{}
	)

	err := birdDB.Create(in.BirdName, in.BirdType, in.Description, in.PicUrl)
	if err != nil {
		return nil, errcode.ErrBirdCreate
	}
	return resp, nil
}
