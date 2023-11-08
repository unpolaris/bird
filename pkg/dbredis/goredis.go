package dbredis

import goredis "github.com/redis/go-redis/v9"

// Config 缓存相关配置
type Config struct {
	Address     string // 地址
	Password    string // 密码
	DB          int    // 数据库
	MaxIdle     int    // 最大空闲连接数
	MaxActive   int    // 最大活跃连接数
	IdleTimeout int    // 空闲连接超时时间
}

func MustNewRedisClient(c *Config) goredis.UniversalClient {
	cli := goredis.NewClient(&goredis.Options{
		Addr:       c.Address,
		Password:   c.Password,
		DB:         c.DB,
		PoolSize:   50,
		MaxRetries: 10,
	})
	return cli
}
