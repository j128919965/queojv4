package config

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/zrpc"
)


type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}

	Redis redis.RedisConf
	ProblemClient zrpc.RpcClientConf
}
