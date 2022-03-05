package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"it-ku/app/usercenter/cmd/rpc/internal/config"
	"it-ku/app/usercenter/model"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	UserModel   model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.NewRedis(c.Redis.Host, c.Redis.Type),
		UserModel:   model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
