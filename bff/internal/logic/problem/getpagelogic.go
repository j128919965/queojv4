package problem

import (
	"context"
	"queoj/service/problem/problem"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetPageLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetPageLogic {
	return GetPageLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPageLogic) GetPage() (resp *types.ProblemPage, err error) {
	page, err := l.svcCtx.ProblemClient.GetAllProblem(l.ctx,&problem.Empty{})
	if err != nil {
		return nil, err
	}

	var problems []*types.ProblemSynopsis
	for _,v := range page.Problems {
		problems = append(problems,&types.ProblemSynopsis{
			Id:    v.Id,
			Name:  v.Name,
			Point: v.Point,
			Level: v.Level,
			Tags:  v.Tags,
		})
	}

	return &types.ProblemPage{
		TotalCount: page.TotalCount,
		Problems:   problems,
	},nil
}
