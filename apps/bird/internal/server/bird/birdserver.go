// Code generated by goctl. DO NOT EDIT.
// Source: bird.proto

package server

import (
	"context"

	"birdProtection/apps/bird/internal/logic/bird"
	"birdProtection/apps/bird/internal/svc"
	"birdProtection/apps/bird/pb/birdservice"
)

type BirdServer struct {
	svcCtx *svc.ServiceContext
	birdservice.UnimplementedBirdServer
}

func NewBirdServer(svcCtx *svc.ServiceContext) *BirdServer {
	return &BirdServer{
		svcCtx: svcCtx,
	}
}

func (s *BirdServer) BirdAdd(ctx context.Context, in *birdservice.BirdAddReq) (*birdservice.BirdAddResp, error) {
	l := birdlogic.NewBirdAddLogic(ctx, s.svcCtx)
	return l.BirdAdd(in)
}

func (s *BirdServer) BirdUpdate(ctx context.Context, in *birdservice.BirdUpdateReq) (*birdservice.BirdUpdateResp, error) {
	l := birdlogic.NewBirdUpdateLogic(ctx, s.svcCtx)
	return l.BirdUpdate(in)
}

func (s *BirdServer) BirdList(ctx context.Context, in *birdservice.BirdListReq) (*birdservice.BirdListResp, error) {
	l := birdlogic.NewBirdListLogic(ctx, s.svcCtx)
	return l.BirdList(in)
}

func (s *BirdServer) BirdInfo(ctx context.Context, in *birdservice.BirdInfoReq) (*birdservice.BirdInfoResp, error) {
	l := birdlogic.NewBirdInfoLogic(ctx, s.svcCtx)
	return l.BirdInfo(in)
}

func (s *BirdServer) BirdSearch(ctx context.Context, in *birdservice.BirdSearchReq) (*birdservice.BirdSearchResp, error) {
	l := birdlogic.NewBirdSearchLogic(ctx, s.svcCtx)
	return l.BirdSearch(in)
}

func (s *BirdServer) BirdDelete(ctx context.Context, in *birdservice.BirdDeleteReq) (*birdservice.BirdDeleteResp, error) {
	l := birdlogic.NewBirdDeleteLogic(ctx, s.svcCtx)
	return l.BirdDelete(in)
}
