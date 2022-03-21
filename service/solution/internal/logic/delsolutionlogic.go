package logic

import (
	"context"

	"queoj/service/solution/internal/svc"
	"queoj/service/solution/solution"

	"github.com/tal-tech/go-zero/core/logx"
)

type DelSolutionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelSolutionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSolutionLogic {
	return &DelSolutionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelSolutionLogic) DelSolution(in *solution.SolutionByIdReq) (*solution.Empty, error) {
	err := l.svcCtx.Db.Exec("delete from solutions where id = ?", in.Id).Error
	if err != nil{
		return nil, err
	}
	return &solution.Empty{}, nil
}
