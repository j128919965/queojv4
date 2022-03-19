package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/j128919965/gopkg/errors"
	"github.com/j128919965/gopkg/security"
	"github.com/tal-tech/go-zero/core/bloom"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/zrpc"
	"gorm.io/gorm"
	"queoj/service/email/emailclient"
	"queoj/service/message/messageclient"
	"queoj/service/user/internal/config"
	"queoj/service/user/internal/model"
	"strconv"
)

type ServiceContext struct {
	ctx                  context.Context
	stop                 func()
	TokenGenerator       *security.TokenGenerator
	TokenParser          *security.TokenParser
	Config               config.Config
	Db                   *gorm.DB
	Redis                *redis.Redis
	CoinsAndPointsSyncer *CoinsAndPointsSyncer
	EmailClient          emailclient.Email
	MessageClient        messageclient.Message
	Filter               *bloom.Filter
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := model.NewMysql(c.Mysql.DataSource)
	if err != nil {
		panic(err.Error())
	}
	ctx, stop := context.WithCancel(context.Background())
	return &ServiceContext{
		ctx:                  ctx,
		stop:                 stop,
		Config:               c,
		Db:                   db,
		Redis:                redis.New(c.Redis.Host),
		CoinsAndPointsSyncer: NewCoinsAndPointsSyncer(ctx, db),
		EmailClient:          emailclient.NewEmail(zrpc.MustNewClient(c.EmailClient)),
		MessageClient:        messageclient.NewMessage(zrpc.MustNewClient(c.MessageClient)),
		TokenGenerator:       security.NewTokenGenerator(c.JwtAuth.AcKey, c.JwtAuth.RfKey),
		TokenParser:          security.NewTokenParser(c.JwtAuth.AcKey, c.JwtAuth.RfKey),
	}
}

func (s *ServiceContext) Stop() {
	s.stop()
}

func (s *ServiceContext) GetUserInfo(id uint64) (*model.UserInfo, error) {
	key := fmt.Sprintf("u:i:%d", id)
	str, err := s.Redis.Get(key)
	if err != nil {
		return nil, err
	}
	if str == "emp" {
		return nil, errors.NotFound()
	}

	info := model.UserInfo{}
	if str == "" {
		err := s.Db.Where(&model.UserInfo{ID: id}).First(&info).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				s.Redis.Setex(key, "emp", 60*2)
				return nil, errors.NotFound()
			} else {
				return nil, err
			}
		}
		bytes, _ := json.Marshal(&info)
		s.Redis.Setex(key, string(bytes), 60*20)
		return &info, nil
	}

	err = json.Unmarshal([]byte(str), &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (s *ServiceContext) RemoveUserInfoCache(id uint64) {
	key := fmt.Sprintf("u:i:%d", id)
	s.Redis.Del(key)
}

func (s *ServiceContext) GetUserAccount(id uint64) (*model.UserAccount, error) {

	key := fmt.Sprintf("u:a:%d", id)
	str, err := s.Redis.Get(key)
	if err != nil {
		return nil, err
	}
	if str == "emp" {
		return nil, errors.NotFound()
	}

	account := model.UserAccount{}
	if str == "" {
		err := s.Db.Where(&model.UserAccount{ID: id}).First(&account).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				s.Redis.Setex(key, "emp", 60*2)
				return nil, errors.NotFound()
			} else {
				return nil, err
			}
		}
		bytes, _ := json.Marshal(&account)
		s.Redis.Setex(key, string(bytes), 60*20)
		return &account, nil
	}

	err = json.Unmarshal([]byte(str), &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *ServiceContext) RemoveUserAccountCache(id uint64) {
	key := fmt.Sprintf("u:a:%d", id)
	s.Redis.Del(key)
}

func (s *ServiceContext) GetUidByEmail(email string) (uint64, error) {
	key := fmt.Sprintf("u:map:e:%s", email)
	str, err := s.Redis.Get(key)
	if err != nil {
		return 0, err
	}
	if str == "emp" {
		return 0, errors.NotFound()
	}

	user := model.User{}
	if str == "" {
		err := s.Db.Where(&model.User{Email: email}).First(&user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				s.Redis.Setex(key, "emp", 60*2)
				return 0, errors.NotFound()
			} else {
				return 0, err
			}
		}
		s.Redis.Setex(key, strconv.FormatUint(user.ID, 10), 60*20)
		return user.ID, nil
	}

	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}
