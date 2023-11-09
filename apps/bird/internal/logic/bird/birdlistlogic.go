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
	cacheKey := fmt.Sprintf("%s:%v", gconst.BirdListPrefix, in.Page)
	if in.Page == 1 {
		cacheStr, err := l.svcCtx.CacheDB.Get(l.ctx, cacheKey).Result()
		if err != nil {
			if !errcode.IsRedisRecordNotFound(err) {
				return nil, err
			}
		} else {
			err = sonic.UnmarshalString(cacheStr, &resp)
			return resp, nil
		}
	}

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
	if in.Page == 1 {
		cacheStr, _ := sonic.MarshalString(resp)
		err = l.svcCtx.CacheDB.Set(l.ctx, cacheKey, cacheStr, gconst.BirdListExpire).Err()
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}
