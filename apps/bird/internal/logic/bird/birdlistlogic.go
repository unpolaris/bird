package birdlogic

import (
	"birdProtection/model"
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
	var (
		birdDB = model.NewBirdModel(l.svcCtx.BirdDB, l.ctx)
		resp   = &birdservice.BirdListResp{}
	)
	list, err := birdDB.List(int(in.Page), int(in.PageSize))
	if err != nil {
		return nil, err
	}
	for _, b := range list {
		resp.List = append(resp.List, &birdservice.BirdListData{
			BirdID:      b.Id,
			BirdName:    b.BirdName,
			BirdType:    b.BirdType,
			Description: b.Description,
			PicUrl:      b.PicUrl,
		})
	}

	return resp, nil
}
