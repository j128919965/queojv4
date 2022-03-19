package config

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/zrpc"
)


type Config struct {
	zrpc.RpcServerConf
	Redis redis.RedisConf
	Smtp struct{
		Host string
		User string
		Password string
		Port int
	}
	Threads int
}
