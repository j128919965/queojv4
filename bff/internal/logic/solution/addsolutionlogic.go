package solution

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"queoj/service/solution/solutionclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddSolutionLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSolutionLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddSolutionLogic {
	return AddSolutionLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSolutionLogic) AddSolution(req types.SolutionAddReq) error {
	payLoad := l.ctx.Value("payload").(*security.PayLoad)
	userId := payLoad.UserId
	isTeacher := false
	if payLoad.Role == 2 {
		isTeacher = true
	}
	addReq := solutionclient.SolutionAddReq{
		Pid:      req.Pid,
		Uid:      userId,
		Nickname: req.Nickname,
		Title:    req.Title,
		Content:  req.Content,
		IsTeacher: isTeacher,
	}
	_, err := l.svcCtx.SolutionClient.AddSolution(l.ctx, &addReq)
	return err
}
