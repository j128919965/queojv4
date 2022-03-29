package svc

import (
	"context"
	"encoding/json"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"gopkg.in/gomail.v2"
	"math/rand"
	"queoj/service/email/internal/config"
	"time"
)

var redisMailKey string = "mailBuf"

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
	ctx context.Context
	stop func()
}

type SendMailInfo struct {
	Address string
	Content string
	Subject string
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx , stop := context.WithCancel(context.Background())
	return &ServiceContext{
		Config: c,
		Redis:  redis.New(c.Redis.Host,redis.WithPass(c.Redis.Pass)),
		ctx: ctx,
		stop: stop,
	}
}

func (srv *ServiceContext) Start() {
	for i:=0 ; i <srv.Config.Threads ; i++ {
		go func(srv *ServiceContext,name int) {
			for {
				select {
				case <-srv.ctx.Done():
					logx.Infof("sender worker %d exit .. ",name)
					return
				default:
					rpop, err := srv.Redis.Rpop(redisMailKey)
					if err != nil {
						if err!=redis.Nil {
							logx.Errorf("get from redis error :%v",err)
						}
						time.Sleep(time.Duration(rand.Intn(100) + 4950) * time.Millisecond)
						break
					}
					info := SendMailInfo{}
					err = json.Unmarshal([]byte(rpop), &info)
					if err != nil {
						logx.Errorf("unmarshal failed : %v",err)
						break
					}
					srv.SendMailSync(&info)
				}
			}	

		}(srv,1)
	}
}

func (srv *ServiceContext) SendMailSync(info *SendMailInfo) error {
	message := gomail.NewMessage()
	message.SetBody("text/html",info.Content)
	message.SetHeader("To",info.Address)
	message.SetHeader("Subject",info.Subject)
	message.SetHeader("From",message.FormatAddress(srv.Config.Smtp.User, "QueOJ"))

	d := gomail.NewDialer(srv.Config.Smtp.Host, srv.Config.Smtp.Port, srv.Config.Smtp.User, srv.Config.Smtp.Password)
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(message)
	if err != nil {
		logx.Errorf("send mail failed : %v")
		return err
	}
	logx.Infof("send mail success , Mail { To : %s, Subject: %s}",info.Address,info.Subject)
	return nil
}

func (srv *ServiceContext) SendMailAsync(info *SendMailInfo) error {
	marshal, err := json.Marshal(info)
	if err != nil {
		return err
	}
	val := string(marshal)
	_, err = srv.Redis.Lpush(redisMailKey, val)
	if err != nil {
		logx.Errorf("add mail to redis failed : %v",err.Error())
		return err
	}
	logx.Infof("add mail to redis success : Mail { To : %s, Subject: %s}",info.Address,info.Subject)
	return nil
}