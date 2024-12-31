package main

import (
	"context"
	"flag"
	"fmt"
	"go-api/internal/jobCron"
	"go-api/pkg/app"
	"go-api/pkg/appMiddleware"
	"go-api/pkg/logs/xcode"

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
		app.WithCheckTokenHandleSMOption(middleware.CheckTokenHandle(ctx)),
	}

	jobCron.NewJobCron(context.Background(), ctx).Run()

	if c.IsDebug {
		opt = append(opt, app.WithDebugOption())
	}
	app.NewServerMiddleware(server, opt...).ApiUseMiddleware()
	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandler(xcode.ErrHandler)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
