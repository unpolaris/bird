package errcode

import (
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func IsRedisRecordNotFound(err error) bool {
	return errors.Is(err, redis.Nil)
}
