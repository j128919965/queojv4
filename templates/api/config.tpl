package config

import (
    "github.com/tal-tech/go-zero/core/stores/redis"
    {{.authImport}}
)

type Config struct {
	rest.RestConf
	Redis redis.RedisConf
	{{.auth}}
}
