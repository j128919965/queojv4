package solution

import (
	"context"
	"queoj/service/solution/solutionclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AdminEditSolutionLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminEditSolutionLogic(ctx context.Context, svcCtx *svc.ServiceContext) AdminEditSolutionLogic {
	return AdminEditSolutionLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminEditSolutionLogic) AdminEditSolution(req types.SolutionDetail) error {
	_, err := l.svcCtx.SolutionClient.EditSolution(l.ctx, &solutionclient.SolutionDetail{
		Id:       req.Id,
		Uid:      req.Uid,
		Pid:      req.Pid,
		Time:     req.Time,
		Nickname: req.Nickname,
		Title:    req.Title,
		Summary:  req.Summary,
		Content:  req.Content,
	})
	return err
}
