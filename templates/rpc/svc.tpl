package svc

import (
    {{.imports}}
    "github.com/tal-tech/go-zero/core/stores/redis"
    "gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db *gorm.DB
	Redis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := model.NewMysql(c.Mysql.DataSource)
	if err != nil {
		panic(err.Error())
	}
	return &ServiceContext{
		Config:    c,
		Db:        db,
		Redis:     redis.New(c.Redis.Host),
	}
}
