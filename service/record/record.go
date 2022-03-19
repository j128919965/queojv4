package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"queoj/service/record/internal/config"
	"queoj/service/record/internal/server"
	"queoj/service/record/internal/svc"
	"queoj/service/record/record"
)

var configFile = flag.String("f", "service/record/etc/record.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewRecordServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		record.RegisterRecordServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)

	//addRecord, err := logic.NewAddRecordLogic(context.Background(), ctx).AddRecord(&record.RecordDetail{
	//	Uid:      1,
	//	Time:     uint64(time.Now().Unix()),
	//	Pid:      1,
	//	Language: 1,
	//	Code:     "public class Main{\n    \n    private static boolean b = true;\n\n    public static void main(String[] args){\n        while(b){\n            System.out.println(\"hello worsfdsld\");\n        }\n        System.out.println(\"hello worsfdsld\");\n    }\n    \n}",
	//})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(addRecord)
	//

	s.Start()
}
