package svc

import (
	"github.com/redis/go-redis/v9"
	"go-api/internal/config"
	"go-api/internal/pkg/dbM"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Client
	DbSelect    *dbM.SelectDb
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		DbSelect:    dbM.NewSelectDb(c.SelectDb),
		RedisClient: dbM.RedisConnect(&redis.Options{Addr: c.RedisConf.Addr, Username: c.RedisConf.Username, Password: c.RedisConf.Password, DB: c.RedisConf.DB}),
	}
}
