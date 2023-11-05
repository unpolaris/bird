package bird

import (
	"context"

	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBirUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirUpdateLogic {
	return &BirUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BirUpdateLogic) BirUpdate(req *types.BirUpdateReq) (resp *types.BirUpdateResp, err error) {


	return
}
