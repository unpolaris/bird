package birdlogic

import (
	"birdProtection/model"
	"birdProtection/pkg/errcode"
	"context"

	"birdProtection/apps/bird/internal/svc"
	"birdProtection/apps/bird/pb/birdservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirdSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBirdSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirdSearchLogic {
	return &BirdSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BirdSearchLogic) BirdSearch(in *birdservice.BirdSearchReq) (*birdservice.BirdSearchResp, error) {
	var (
		birdDB = model.NewBirdModel(l.svcCtx.BirdDB, l.ctx)
		resp   = &birdservice.BirdSearchResp{}
	)
	list, err := birdDB.Search(in.BirdName)
	if err != nil {
		return nil, errcode.ErrBirdSearch
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
