package logic

import (
	"context"
	"github.com/j128919965/gopkg/errors"

	"queoj/service/problem/internal/svc"
	"queoj/service/problem/problem"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProblemDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemDetailLogic {
	return &GetProblemDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProblemDetailLogic) GetProblemDetail(in *problem.Integer) (*problem.ProblemDetail, error) {
	if in == nil || in.Value < 1 {
		return nil,errors.New("输入ID有误",400)
	}
	p, tags := l.svcCtx.GetProblem(in.Value)
	return ProblemToProblemDetail(p,tags), nil
}
