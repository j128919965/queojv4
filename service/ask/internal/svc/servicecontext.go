package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"queoj/service/ask/internal/config"
	"queoj/service/ask/internal/model"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := model.NewMysql(c.Mysql.DataSource)
	if err != nil {
		panic(err.Error())
	}
	return &ServiceContext{
		Config: c,
		Db:     db,
		Redis:  redis.New(c.Redis.Host),
	}
}

func (svc *ServiceContext) GetAllAsk() []*model.Ask {
	var asks []*model.Ask
	svc.Db.Select([]string{"id","uid","time","nick_name","title"}).Find(&asks)
	return asks
}

func (svc *ServiceContext) GetAskById(id uint64) (*model.Ask,error){
	var ask model.Ask
	err := svc.Db.Where("id = ?", id).Find(&ask).Error
	if err != nil {
		return nil, err
	}
	return &ask,nil
}

func (svc *ServiceContext) GetRepliesByAsk(askId uint64) []*model.Reply{
	var replies []*model.Reply
	svc.Db.Where("ask_id = ?",askId).Find(&replies)
	return replies
}