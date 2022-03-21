package solution

import (
	"context"
	"queoj/service/solution/solutionclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteSolutionLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSolutionLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteSolutionLogic {
	return DeleteSolutionLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSolutionLogic) DeleteSolution(req types.SolutionByIdReq) error {
	_, err := l.svcCtx.SolutionClient.DelSolution(l.ctx, &solutionclient.SolutionByIdReq{Id: req.Id})
	return err
}
