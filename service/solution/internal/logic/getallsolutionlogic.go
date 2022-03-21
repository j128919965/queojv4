package logic

import (
	"context"

	"queoj/service/solution/internal/svc"
	"queoj/service/solution/solution"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllSolutionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllSolutionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllSolutionLogic {
	return &GetAllSolutionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllSolutionLogic) GetAllSolution(in *solution.Empty) (*solution.SolutionList, error) {
	solutions, err := l.svcCtx.GetAllSolutions()
	if err != nil {
		return nil, err
	}
	var ret []*solution.SolutionSummary
	for _, m := range solutions {
		ret = append(ret,ModelToInfo(m))
	}
	return &solution.SolutionList{Solutions: ret}, nil
}
