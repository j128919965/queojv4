package solution

import (
	"context"
	"github.com/j128919965/gopkg/errors"
	"github.com/j128919965/gopkg/security"
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
	payLoad := l.ctx.Value("payload").(*security.PayLoad)
	if payLoad.Role < 2 {
		return errors.New("权限不足",400)
	}
	_, err := l.svcCtx.SolutionClient.DelSolution(l.ctx, &solutionclient.SolutionByIdReq{Id: req.Id})
	return err
}
