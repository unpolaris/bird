package config

import (
	"birdProtection/pkg/dbmysql"
	"birdProtection/pkg/dbredis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	BirdDB dbmysql.Config
	Cache  dbredis.Config
}
