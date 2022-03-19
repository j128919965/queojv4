package logic

import (
	"context"
	"encoding/json"
	"queoj/service/problem/internal/model"

	"queoj/service/problem/internal/svc"
	"queoj/service/problem/problem"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllProblemLogic {
	return &GetAllProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllProblemLogic) GetAllProblem(_ *problem.Empty) (*problem.ProblemPage, error) {
	problems,ids := l.svcCtx.GetAllProblems()
	page := problem.ProblemPage{TotalCount: int32(len(problems))}
	for i := 0; i < len(problems); i++ {
		page.Problems = append(page.Problems , ProblemToProblemSynopsis(problems[i],ids[i]))
	}
	return &page, nil
}

func ProblemToProblemSynopsis (p *model.Problem , tags []int32) *problem.ProblemSynopsis {
	str,_ := json.Marshal(tags)
	s := problem.ProblemSynopsis{
		Id:    p.Id,
		Name:  p.Name,
		Point: p.Point,
		Level: p.Level,
		Tags:  string(str),
	}
	return &s
}


func ProblemToProblemDetail (p *model.Problem , tags []int32) *problem.ProblemDetail {
	str,_ := json.Marshal(tags)
	s := problem.ProblemDetail{
		Id:          p.Id,
		Name:        p.Name,
		Point:       p.Point,
		Level:       p.Level,
		Description: p.Description,
		Tags:        string(str),
		TimeLimit:   p.TimeLimit,
		SpaceLimit:  p.SpaceLimit,
	}
	return &s
}
