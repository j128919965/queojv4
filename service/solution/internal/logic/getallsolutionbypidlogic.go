package logic

import (
	"context"
	"queoj/service/solution/internal/model"

	"queoj/service/solution/internal/svc"
	"queoj/service/solution/solution"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllSolutionByPidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllSolutionByPidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllSolutionByPidLogic {
	return &GetAllSolutionByPidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllSolutionByPidLogic) GetAllSolutionByPid(in *solution.AllSolutionByPidReq) (*solution.SolutionList, error) {
	solutions, err := l.svcCtx.GetAllSolutionsByPid(in.Pid)
	if err != nil {
		return nil, err
	}
	var ret []*solution.SolutionSummary
	for _, m := range solutions {
		ret = append(ret,ModelToSummary(m))
	}
	return &solution.SolutionList{Solutions: ret}, nil
}

func ModelToSummary(s *model.Solution) *solution.SolutionSummary {
	return &solution.SolutionSummary{
		Id:       s.Id,
		Time:     s.Time,
		Nickname: s.Nickname,
		Title:    s.Title,
		Summary:  s.Summary,
	}
}

func ModelToInfo(s *model.Solution) *solution.SolutionSummary {
	return &solution.SolutionSummary{
		Id:       s.Id,
		Time:     s.Time,
		Nickname: s.Nickname,
		Title:    s.Title,
	}
}