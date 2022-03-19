package svc

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/zrpc"
	"gorm.io/gorm"
	"queoj/service/problem/problemclient"
	"queoj/service/record/internal/config"
	"queoj/service/record/internal/model"
	record_status "queoj/service/record/internal/svc/record-status"
)



type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	Redis  *redis.Redis
	ProblemClient problemclient.Problem
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
		ProblemClient: problemclient.NewProblem(zrpc.MustNewClient(c.ProblemClient)),
	}
}

func (svc *ServiceContext) GetRecordById(id uint64) (*model.Record,error) {
	var record *model.Record
	err := svc.Db.Where("id = ?", id).Find(&record).Error
	if err != nil {
		return nil,err
	}
	return record,nil
}

func (svc *ServiceContext) InsertRecord(record *model.Record) error {
	return svc.Db.Create(record).Error
}

func (svc *ServiceContext) GetRecordByUserId(userId uint64) ([]*model.Record,error) {
	var record []*model.Record
	err := svc.Db.Where("uid = ?", userId).Find(&record).Error
	if err != nil {
		return nil,err
	}
	return record,nil
}

func (svc *ServiceContext) UpdateRecordStatus(id uint64,status uint32)  {
	svc.Db.Model(&model.Record{}).Where("id = ?",id).Update("status",status)
}

func (svc *ServiceContext) UpdateRecordResult(id ,time ,space uint64)  {
	svc.Db.Model(&model.Record{}).Where("id = ?",id).Updates(&model.Record{Status: record_status.Accept,TimeUsed: time,SpaceUsed: space})
}

const JudgeUrl = `http://localhost:5050/run`

func (svc *ServiceContext) SubmitRecord (record *model.Record) {
	logx.Info("start judge {%d}",record.Id)
	switch record.Language {
	case 1: svc.submitJava(record)
	case 2: svc.submitCpp(record)
	case 3: svc.submitGo(record)
	default:
		panic("暂不支持该语言！")
	}

}
