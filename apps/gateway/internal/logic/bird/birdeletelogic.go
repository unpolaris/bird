package bird

import (
	"context"

	"birdProtection/apps/gateway/internal/svc"
	"birdProtection/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BirDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBirDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BirDeleteLogic {
	return &BirDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BirDeleteLogic) BirDelete(req *types.BirDeleteReq) (resp *types.BirDeleteReq, err error) {
	// todo: add your logic here and delete this line

	return
}
