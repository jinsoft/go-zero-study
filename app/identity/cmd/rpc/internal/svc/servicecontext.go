package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"it-ku/app/identity/cmd/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.New(c.Redis.Host),
	}
}
