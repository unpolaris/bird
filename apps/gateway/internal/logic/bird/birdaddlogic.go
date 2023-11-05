package bird

import (
	"context"

	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirdAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBirdAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirdAddLogic {
	return &BirdAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BirdAddLogic) BirdAdd(req *types.BirdAddReq) (resp *types.BirdAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
