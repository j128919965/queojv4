package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"queoj/service/solution/internal/config"
	"queoj/service/solution/internal/model"
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
		Redis:  redis.New(c.Redis.Host,redis.WithPass(c.Redis.Pass)),
	}
}

func (svc *ServiceContext) GetSolutionById(id uint64) (*model.Solution,error){
	var solution *model.Solution
	err := svc.Db.Where("id = ?", id).Find(&solution).Error
	if err != nil {
		return nil,err
	}
	return solution,nil
}

func (svc *ServiceContext) GetAllSolutionsByPid(id int32) ([]*model.Solution,error) {
	var solutions []*model.Solution
	err := svc.Db.Where("pid = ?", id).Find(&solutions).Error
	if err != nil {
		return nil,err
	}
	return solutions,nil
}

func (svc *ServiceContext) GetAllSolutions() ([]*model.Solution,error) {
	var solutions []*model.Solution
	err := svc.Db.Find(&solutions).Error
	if err != nil {
		return nil,err
	}
	return solutions,nil
}

