package main

import (
	"flag"
	"fmt"

	"queoj/bff/internal/config"
	"queoj/bff/internal/handler"
	"queoj/bff/internal/svc"

	"github.com/j128919965/gopkg/web/middleware"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "bff/etc/queoj-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf,rest.WithCors())
	defer server.Stop()

	server.Use(middleware.ErrorHandler)
	server.Use(middleware.AuthHandler)



	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
