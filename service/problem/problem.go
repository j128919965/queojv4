package main

import (
	"context"
	"flag"
	"fmt"
	"queoj/service/problem/internal/logic"

	"queoj/service/problem/internal/config"
	"queoj/service/problem/internal/server"
	"queoj/service/problem/internal/svc"
	"queoj/service/problem/problem"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "service/problem/etc/problem.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		problem.RegisterProblemServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	allProblem, _ := logic.NewGetProblemsByTagsLogic(context.Background(), ctx).GetProblemsByTags(&problem.ProblemsByTagReq{Tags: []int32{4}})
	fmt.Println(allProblem)
	s.Start()
}
