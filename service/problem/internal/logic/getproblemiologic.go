package logic

import (
	"context"
	"github.com/j128919965/gopkg/errors"

	"queoj/service/problem/internal/svc"
	"queoj/service/problem/problem"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProblemIOLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemIOLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemIOLogic {
	return &GetProblemIOLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProblemIOLogic) GetProblemIO(in *problem.Integer) (*problem.ProblemIO, error) {
	if in == nil || in.Value < 1 {
		return nil,errors.New("输入ID有误",400)
	}
	io := l.svcCtx.GetProblemIO(in.Value)
	return &problem.ProblemIO{
		Id:     io.Id,
		Pid:    io.Pid,
		InTxt:  io.InTxt,
		OutTxt: io.OutTxt,
	}, nil
}
