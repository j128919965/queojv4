package config

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Redis redis.RedisConf
	Auth  struct {
		AccessSecret string
		AccessExpire int64
	}
	UserClient zrpc.RpcClientConf
	ProblemClient zrpc.RpcClientConf
	SolutionClient zrpc.RpcClientConf
	MessageClient zrpc.RpcClientConf
	RecordClient zrpc.RpcClientConf
	AskClient zrpc.RpcClientConf
}
