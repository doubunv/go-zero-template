package main

import (
	"flag"
	"fmt"
	"go-api/internal/pkg/app"
	"go-api/internal/pkg/appMiddleware"
	"go-api/internal/pkg/logs/xcode"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-api/internal/config"
	"go-api/internal/handler"
	"go-api/internal/middleware"
	"go-api/internal/svc"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	server := rest.MustNewServer(c.RestConf, rest.WithNotAllowedHandler(appMiddleware.NewCorsMiddleware().Handler()))
	defer server.Stop()

	app.InitAppServer()
	ctx := svc.NewServiceContext(c)
	var opt = []app.SMOption{
		app.WithWhiteHeaderPathSMOption(middleware.WhiteHeaderPath()),
		app.WithCheckTokenHandleSMOption(middleware.CheckTokenHandle(c)),
	}

	if c.IsDebug {
		opt = append(opt, app.WithDebugOption())
	}
	app.NewServerMiddleware(server, opt...).ApiUseMiddleware()
	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandler(xcode.ErrHandler)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
