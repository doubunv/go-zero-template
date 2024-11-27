package app

import (
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
	"go-api/pkg/dbM"
	"go-api/pkg/logs"
)

func init() {
}

func logInit() {
	// 设置日志
	writer := logs.NewLogrusWriter(func(logger *logrus.Logger) {
		logger.SetFormatter(&logrus.JSONFormatter{})
	})
	logx.SetWriter(writer)
}

func InitAppServer() {
	logInit()
	logx.DisableStat()
}

type RedisConf struct {
	Addr     string
	Username string
	Password string
	DB       int
}

type ConfigAppServer struct {
	IsDebug   bool                 `json:",optional,default=false"`
	SelectDb  []dbM.SelectDbConfig `json:",optional"`
	RedisConf *RedisConf           `json:",optional"`
}
