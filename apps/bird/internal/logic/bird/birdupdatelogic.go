package birdlogic

import (
	"birdProtection/model"
	"birdProtection/pkg/errcode"
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
	var (
		birdDB = model.NewBirdModel(l.svcCtx.BirdDB, l.ctx)
		resp   = &birdservice.BirdUpdateResp{}
	)
	err := birdDB.Update(in.BirdID, in.BirdName, in.BirdType, in.Description, in.PicUrl)
	if err != nil {
		return nil, errcode.ErrBirdUpdate
	}
	resp.BirdID = in.BirdID
	return resp, nil
}
