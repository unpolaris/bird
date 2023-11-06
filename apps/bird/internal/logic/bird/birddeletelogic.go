package birdlogic

import (
	"birdProtection/model"
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
	var (
		birdDB = model.NewBirdModel(l.svcCtx.BirdDB, l.ctx)
		resp   = &birdservice.BirdDeleteResp{}
	)
	err := birdDB.Delete(in.BirdID)
	if err != nil {
		return nil, err
	}
	resp.BirdID = in.BirdID
	return resp, nil
}
