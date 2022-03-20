package main

import (
	"flag"
	"fmt"
	"queoj/service/user/internal/config"
	"queoj/service/user/internal/server"
	"queoj/service/user/internal/svc"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "service/user/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	defer ctx.Stop()

	srv := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)

	//rank, err := logic.NewGetUserRankLogic(context.Background(),ctx).GetUserRank(&user.RankByUserIdReq{Id: 1})
	//if err != nil {
	//	logx.Error(err)
	//}else {
	//	logx.Info(rank)
	//}


	s.Start()
}
