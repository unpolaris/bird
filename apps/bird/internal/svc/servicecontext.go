package svc

import (
	"birdProtection/apps/bird/internal/config"
	"birdProtection/pkg/dbmysql"
	"birdProtection/pkg/dbredis"
	goredis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	BirdDB  *gorm.DB
	CacheDB goredis.UniversalClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		BirdDB:  dbmysql.MustNewDB(&c.BirdDB),
		CacheDB: dbredis.MustNewRedisClient(&c.Cache),
	}
}
