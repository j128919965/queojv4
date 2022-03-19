package svc

import (
	"encoding/json"
	"fmt"
	"github.com/j128919965/gopkg/errors"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"queoj/service/message/internal/config"
	"queoj/service/message/internal/model"
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

func (s *ServiceContext) GetMessageById(id uint64) (*model.Message, error) {
	key := fmt.Sprintf("msg:id:%d", id)
	str, err := s.Redis.Get(key)
	if err != nil {
		return nil, err
	}
	if str == "emp" {
		return nil, errors.NotFound()
	}

	message := model.Message{}
	if str == "" {
		err := s.Db.Where(&model.Message{Id: id}).First(&message).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				s.Redis.Setex(key, "emp", 60*2)
				return nil, errors.NotFound()
			} else {
				return nil, err
			}
		}
		bytes, err := json.Marshal(&message)
		if err != nil {
			return nil, err
		}
		s.Redis.Setex(key, string(bytes), 60*20)
		return &message, nil
	}

	err = json.Unmarshal([]byte(str), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (s *ServiceContext) GetNoReadMessagesByUser(user uint64, noread bool ,pageSize int, pageFrom *int32, lastId *uint64) ([]model.Message, error) {
	sql := "SELECT id,receiver,`read`,type,title,created_at FROM `messages` WHERE receiver=? "
	var datas []interface{}
	datas = append(datas,user)
	if noread {
		sql += "and `read`=0 "
	}
	sql += "and deleted_at IS null "
	if lastId!=nil {
		sql += "and id > ?"
		datas = append(datas , lastId)
	}
	sql += "order by `read` limit "
	if lastId==nil && pageFrom != nil {
		sql += "? , "
		datas = append(datas,pageFrom)
	}
	sql += "?"
	datas = append(datas,pageSize)

	logx.Infof("sql : %s",sql)
	var msgs []model.Message
	err := s.Db.Raw(sql, datas...).Scan(&msgs).Error
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (s *ServiceContext) ReadMsg(id uint64) error {
	key := fmt.Sprintf("msg:id:%d", id)
	s.Redis.Del(key)
	sql := "update messages set `read` = 1 where id = ?"
	return s.Db.Exec(sql, id).Error
}

func (s *ServiceContext) ReadAllMsg(user uint64) error {
	sql := "update messages set `read` = 1 where receiver = ?"
	return s.Db.Exec(sql, user).Error
}

func (s *ServiceContext) HasNoRead(user uint64) bool {
	sql := "select exists(select id from messages where receiver = ? and `read` = 0 and deleted_at is null)"
	var e int
	err := s.Db.Raw(sql, user).Scan(&e).Error
	if err != nil {
		logx.Error(err)
	}
	return e > 0
}

func (s ServiceContext) GetTotalCount(user uint64 , noread bool) (int,error) {
	sql := "select count(1) from messages where receiver = ? and deleted_at IS null"

	if noread {
		sql += "and `read` = 0 ;"
	}

	var i int
	err := s.Db.Raw(sql, user).Scan(&i).Error
	if err != nil {
		return 0,err
	}
	return i,nil
}