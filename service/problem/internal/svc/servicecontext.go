package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"queoj/service/problem/internal/config"
	"queoj/service/problem/internal/model"
	"strconv"
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

func (svc *ServiceContext) GetAllProblems() (problems []*model.Problem,tags [][]int32) {
	svc.Db.Find(&problems)
	var ids []int32
	for _,v := range problems {
		ids = append(ids,v.Id)
	}
	if len(ids) == 0 {
		return
	}
	tags = make([][]int32,len(ids))
	var pts []*model.ProblemTag
	svc.Db.Where("pid in ?",ids).Find(&pts)
	m := map[int32]int{}
	for i,v := range problems {
		m[v.Id] = i
	}
	for _, v := range pts {
		tags[m[v.Pid]] = append(tags[m[v.Pid]],v.Tid)
	}
	return
}

func (svc *ServiceContext) GetProblemsByTagIds(ids []int32) ([]*model.Problem,[][]int32) {
	sqlPre := `select * from problems where id in (select pid from problem_tags where tid in (`
	sqlSur := `) )`
	var problems []*model.Problem
	if len(ids) == 0 {
		return problems,make([][]int32,0)
	}
	for i,id := range ids {
		sqlPre+=strconv.FormatInt(int64(id),10)
		if i != len(ids)-1 {
			sqlPre+=","
		}
	}
	sqlPre += sqlSur
	svc.Db.Raw(sqlPre).Scan(&problems)

	var tags = make([][]int32,len(problems))
	var pts []*model.ProblemTag

	pids := make([]int32,len(problems))
	for _, problem := range problems {
		pids = append(pids,problem.Id)
	}

	svc.Db.Where("pid in ?",pids).Find(&pts)
	m := map[int32]int{}
	for i,v := range problems {
		m[v.Id] = i
	}
	for _, v := range pts {
		tags[m[v.Pid]] = append(tags[m[v.Pid]],v.Tid)
	}

	return problems,tags
}

func (svc ServiceContext) GetProblemIO(id int32) *model.ProblemIO {
	var io model.ProblemIO
	svc.Db.Where("pid = ?",id).First(&io)
	return &io;
}

func (svc ServiceContext) GetProblem(id int32) (*model.Problem,[]int32 ){
	var io model.Problem
	svc.Db.Where("id = ?",id).First(&io)

	var pts []*model.ProblemTag
	svc.Db.Where("pid = ?",id).Find(&pts)
	var tags []int32
	for _,v := range pts {
		tags = append(tags , v.Tid)
	}
	return &io,tags
}