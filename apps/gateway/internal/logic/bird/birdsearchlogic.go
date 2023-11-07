package bird

import (
	birdClient "birdProtection/apps/bird/client/bird"
	"context"

	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirdSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBirdSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirdSearchLogic {
	return &BirdSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BirdSearchLogic) BirdSearch(req *types.BirdSearchReq) (resp *types.BirListResp, err error) {
	res, err := l.svcCtx.BirdClient.BirdSearch(l.ctx, &birdClient.BirdSearchReq{
		BirdName: req.Keyword,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.BirListResp{}
	for _, b := range res.List {
		resp.List = append(resp.List, types.BirListData{
			BirdId:      b.BirdID,
			BirdName:    b.BirdName,
			BirdType:    b.BirdType,
			Description: b.Description,
			PicUrl:      b.PicUrl,
		})
	}
	return
}
