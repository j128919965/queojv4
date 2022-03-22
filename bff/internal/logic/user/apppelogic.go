package user

import (
	"context"
	"github.com/j128919965/gopkg/errors"
	"github.com/j128919965/gopkg/security"
	"queoj/service/user/userclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AppPeLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppPeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AppPeLogic {
	return AppPeLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppPeLogic) AppPe(req types.ApprovalPrivilegeEscalationReq) error {
	payLoad := l.ctx.Value("payload").(*security.PayLoad)
	if payLoad.Role < 3 {
		return errors.New("权限不足",400)
	}
	_, err := l.svcCtx.UserClient.ApprovalPEReq(l.ctx, &userclient.ApprovalPrivilegeEscalationReq{
		Id:       req.Id,
		Approval: req.Approval,
	})
	return err
}
