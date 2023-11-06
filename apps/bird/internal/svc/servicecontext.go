package svc

import (
	"birdProtection/apps/bird/internal/config"
	"birdProtection/pkg/dbmysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	BirdDB *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		BirdDB: dbmysql.MustNewDB(&c.BirdDB),
	}
}
