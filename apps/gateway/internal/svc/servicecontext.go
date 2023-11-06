package svc

import (
	clientBird "birdProtection/apps/bird/client/bird"
	"birdProtection/apps/gateway/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	BirdClient clientBird.Bird
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		BirdClient: clientBird.NewBird(zrpc.MustNewClient(c.BirdClient)),
	}
}
