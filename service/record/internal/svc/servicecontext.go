package svc

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/zrpc"
	"gorm.io/gorm"
	"queoj/service/problem/problemclient"
	"queoj/service/record/internal/config"
	"queoj/service/record/internal/model"
	record_status "queoj/service/record/internal/svc/record-status"
	"queoj/service/user/userclient"
)

type ServiceContext struct {
	Config        config.Config
	Db            *gorm.DB
	Redis         *redis.Redis
	ProblemClient problemclient.Problem
	UserClient    userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := model.NewMysql(c.Mysql.DataSource)
	if err != nil {
		panic(err.Error())
	}
	return &ServiceContext{
		Config:        c,
		Db:            db,
		Redis:         redis.New(c.Redis.Host),
		ProblemClient: problemclient.NewProblem(zrpc.MustNewClient(c.ProblemClient)),
		UserClient:    userclient.NewUser(zrpc.MustNewClient(c.UserClient)),
	}
}

func (svc *ServiceContext) GetRecordById(id uint64) (*model.Record, error) {
	var record *model.Record
	err := svc.Db.Where("id = ?", id).Find(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (svc *ServiceContext) InsertRecord(record *model.Record) error {
	return svc.Db.Create(record).Error
}

func (svc *ServiceContext) GetRecordByUserId(userId uint64) ([]*model.Record, error) {
	var record []*model.Record
	err := svc.Db.Where("uid = ?", userId).Find(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (svc *ServiceContext) UpdateRecordStatus(id uint64, status uint32) {
	svc.Db.Model(&model.Record{}).Where("id = ?", id).Update("status", status)
}

func (svc *ServiceContext) UpdateRecordResult(id, time, space uint64, record *model.Record) {
	var i int
	err := svc.Db.Raw("select count(1) from records where uid = ? and pid = ? and `status` = 1", record.Uid, record.Pid).Scan(&i).Error
	if err != nil {
		logx.Error(err)
	}else {
		logx.Info(fmt.Sprintf("user { %d } passed problem { %d } for { %d } times.",record.Uid,record.Pid,i))
		// 首次通过
		if i == 0 {
			detail, err := svc.ProblemClient.GetProblemDetail(context.Background(), &problemclient.Integer{Value: record.Pid})
			if err != nil {
				logx.Error(err)
			}else {
				svc.UpdateUserRecordStatistic(record.Uid,detail.Level)
				_, err := svc.UserClient.AddCoinOrPoint(context.Background(),&userclient.AddCoinOrPointReq{
					UserId: record.Uid,
					Points: detail.Point,
					Reason: fmt.Sprintf("首次通过题目 %s , 获得 %d 积分！",detail.Name,detail.Point),
				})
				if err != nil {
					logx.Error(err)
				}
			}
		}
	}

	// 无论如何，一定要执行到这一步
	svc.Db.Model(&model.Record{}).Where("id = ?", id).Updates(&model.Record{Status: record_status.Accept, TimeUsed: time, SpaceUsed: space})
}

func (svc *ServiceContext) UpdateUserRecordStatistic(userId uint64,level int32)  {
	var cnt int32
	err := svc.Db.Raw("select count(1) from user_success_statistics where uid = ?", userId).Scan(&cnt).Error
	if err != nil {
		logx.Error(err)
		return
	}
	if cnt == 0 {
		svc.Db.Create(&model.UserSuccessStatistic{Uid: userId})
	}
	switch level {
	case 1: svc.Db.Model(&model.UserSuccessStatistic{}).Where("uid = ?",userId).Update("easy",gorm.Expr("easy + 1"))
	case 2: svc.Db.Model(&model.UserSuccessStatistic{}).Where("uid = ?",userId).Update("medium",gorm.Expr("medium + 1"))
	case 3: svc.Db.Model(&model.UserSuccessStatistic{}).Where("uid = ?",userId).Update("hard",gorm.Expr("hard + 1"))
	}
}

const JudgeUrl = `http://localhost:5555/run`

func (svc *ServiceContext) SubmitRecord(record *model.Record) {
	logx.Info("start judge {%d}", record.Id)
	switch record.Language {
	case 1:
		svc.submitJava(record)
	case 2:
		svc.submitCpp(record)
	case 3:
		svc.submitGo(record)
	default:
		panic("暂不支持该语言！")
	}
}
