package solution

import (
	"context"
	"queoj/service/solution/solutionclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllSolutionLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllSolutionLogic(ctx context.Context, svcCtx *svc.ServiceContext) AllSolutionLogic {
	return AllSolutionLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllSolutionLogic) AllSolution() (resp *types.SolutionList, err error) {
	ss, err := l.svcCtx.SolutionClient.GetAllSolution(l.ctx, &solutionclient.Empty{})
	if err != nil {
		return nil, err
	}

	var ret []*types.SolutionSummary

	for _, s := range ss.GetSolutions() {
		ret = append(ret, &types.SolutionSummary{
			Id:       s.Id,
			Time:     s.Time,
			Nickname: s.Nickname,
			Title:    s.Title,
			Pid:      s.Pid,
			IsTeacher: s.IsTeacher,
		})
	}

	return &types.SolutionList{Solutions: ret}, nil
}
