package main

import (
	"flag"
	"fmt"

	"bigdata/core/internal/config"
	"bigdata/core/internal/handler"
	"bigdata/core/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/core-api.yaml", "the config file")

//var configFile = flag.String("f", "core-api.yaml", "the config file")

func main() {
	flag.Parse()

	fmt.Printf("configFile:", *configFile)
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
