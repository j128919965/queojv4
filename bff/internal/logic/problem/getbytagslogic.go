package problem

import (
	"context"
	"queoj/service/problem/problem"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetByTagsLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetByTagsLogic {
	return GetByTagsLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByTagsLogic) GetByTags(req types.ProblemsByTagReq) (resp *types.ProblemPage, err error) {
	page, err := l.svcCtx.ProblemClient.GetProblemsByTags(l.ctx, &problem.ProblemsByTagReq{Tags: req.Tags})
	if err!=nil {
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

