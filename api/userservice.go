package main

import (
	"flag"
	"fmt"

	"github.com/zwtesttt/starcloud/api/internal/config"
	"github.com/zwtesttt/starcloud/api/internal/handler"
	"github.com/zwtesttt/starcloud/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/userservice.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	fmt.Println("configFile:", c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	fmt.Println("config:", c.GroupRpc)
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
