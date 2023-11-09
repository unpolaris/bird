package birdlogic

import (
	"birdProtection/model"
	"birdProtection/pkg/errcode"
	"birdProtection/pkg/gconst"
	"context"
	"fmt"
	"github.com/bytedance/sonic"

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
	cacheKey := fmt.Sprintf("%s:%v", gconst.BirdListPrefix, in.BirdID)
	cacheStr, err := l.svcCtx.CacheDB.Get(l.ctx, cacheKey).Result()
	if err != nil {
		if !errcode.IsRedisRecordNotFound(err) {
			return nil, err
		}
	} else {
		err = sonic.UnmarshalString(cacheStr, &resp)
		return resp, nil
	}

	data, err := birdDB.Get(in.BirdID)
	if err != nil {
		return nil, err
	}
	if data == nil {
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
	cacheStr, err = sonic.MarshalString(resp)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.CacheDB.Set(l.ctx, cacheKey, cacheStr, gconst.BirdInfoExpire).Err()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
