package birdlogic

import (
	"birdProtection/model"
	"birdProtection/pkg/gconst"
	"context"
	"fmt"

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
	cacheKey := fmt.Sprintf("%s:%v", gconst.BirdInfoPrefix, in.BirdID)
	err := l.svcCtx.CacheDB.Get(l.ctx, cacheKey).Scan(&resp)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		return resp, nil
	}

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
	err = l.svcCtx.CacheDB.Set(l.ctx, cacheKey, resp, gconst.BirdInfoExpire).Err()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
