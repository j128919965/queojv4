package problem

import (
	"context"
	"queoj/service/problem/problem"
	"strconv"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProblemByIdLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProblemByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProblemByIdLogic {
	return GetProblemByIdLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProblemByIdLogic) GetProblemById(req types.Integer) (resp *types.ProblemDetail, err error) {
	_, err = l.svcCtx.Redis.Zincrby("pro:times",1,strconv.FormatInt(int64(req.Value),10))
	if err != nil {
		return nil, err
	}
	detail, err := l.svcCtx.ProblemClient.GetProblemDetail(l.ctx, &problem.Integer{Value: req.Value})
	if err != nil {
		return nil, err
	}
	return &types.ProblemDetail{
		Id:          detail.Id,
		Name:        detail.Name,
		Point:       detail.Point,
		Level:       detail.Level,
		Tags:        detail.Tags,
		Description: detail.Description,
		TimeLimit:   detail.TimeLimit,
		SpaceLimit:  detail.SpaceLimit,
	},nil
}
