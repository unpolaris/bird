package bird

import (
	"context"

	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBirGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirGetLogic {
	return &BirGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BirGetLogic) BirGet(req *types.BirGetReq) (resp *types.BirGetResp, err error) {
	// todo: add your logic here and delete this line

	return
}
