package config

import (
	"birdProtection/pkg/dbmysql"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	BirdDB dbmysql.Config
}
