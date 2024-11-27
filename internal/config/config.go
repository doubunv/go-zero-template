package config

import (
	"github.com/zeromicro/go-zero/rest"
	"go-api/pkg/app"
)

type Config struct {
	rest.RestConf
	app.ConfigAppServer
}
