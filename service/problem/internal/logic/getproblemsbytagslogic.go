package logic

import (
	"context"

	"queoj/service/problem/internal/svc"
	"queoj/service/problem/problem"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProblemsByTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemsByTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemsByTagsLogic {
	return &GetProblemsByTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProblemsByTagsLogic) GetProblemsByTags(in *problem.ProblemsByTagReq) (*problem.ProblemPage, error) {
	problems, ids := l.svcCtx.GetProblemsByTagIds(in.Tags)
	page := problem.ProblemPage{TotalCount: int32(len(problems))}
	for i := 0; i < len(problems); i++ {
		page.Problems = append(page.Problems , ProblemToProblemSynopsis(problems[i],ids[i]))
	}
	return &page, nil
}
