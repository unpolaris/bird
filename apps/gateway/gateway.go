package main

import (
	"birdProtection/pkg/xhttp"
	"flag"
	"fmt"

	"birdProtection/apps/gateway/internal/config"
	"birdProtection/apps/gateway/internal/handler"
	"birdProtection/apps/gateway/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithNotAllowedHandler(xhttp.NewNotAllowedMiddleware().Handler()))
	ctx := svc.NewServiceContext(c)
	server.Use(xhttp.NewLogMiddleware().Handle)
	server.Use(xhttp.NewCorsMiddleware().Handle)
	server.Use(xhttp.NewRecoverMiddleware().Handle)

	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
