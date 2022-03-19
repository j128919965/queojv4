package svc

import (
    "github.com/tal-tech/go-zero/core/stores/redis"
	{{.configImport}}
)

type ServiceContext struct {
	Config {{.config}}
	Redis *redis.Redis
	{{.middleware}}
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis: redis.New(c.Redis.Host),
		{{.middlewareAssignment}}
	}
}
