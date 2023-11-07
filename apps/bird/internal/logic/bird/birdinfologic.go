package birdlogic

import (
	"birdProtection/model"
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
	var (
		birdDB = model.NewBirdModel(l.svcCtx.BirdDB, l.ctx)
		resp   = &birdservice.BirdInfoResp{}
	)
	data, err := birdDB.Get(in.BirdID)
	if err != nil {
		return nil, err
	}
	if data != nil {
		return nil, err
	}
	resp = &birdservice.BirdInfoResp{
		BirdID:      data.Id,
		BirdName:    data.BirdName,
		BirdType:    data.BirdType,
		Description: data.Description,
		PicUrl:      data.PicUrl,
		CreateTime:  data.CreateTime.Unix(),
		UpdateTime:  data.UpdateTime.Unix(),
	}
	return resp, nil
}
