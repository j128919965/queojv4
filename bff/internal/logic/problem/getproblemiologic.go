package problem

import (
	"context"
	"queoj/service/problem/problemclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProblemIOLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProblemIOLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProblemIOLogic {
	return GetProblemIOLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProblemIOLogic) GetProblemIO(req types.Integer) (resp *types.ProblemIO, err error) {
	io, err := l.svcCtx.ProblemClient.GetProblemIO(l.ctx, &problemclient.Integer{Value: req.Value})
	if err != nil {
		return nil, err
	}
	return &types.ProblemIO{
		InText:  io.InTxt,
		OutText: io.OutTxt,
	},nil
}
