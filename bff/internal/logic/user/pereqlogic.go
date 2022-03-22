package user

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"queoj/service/user/userclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PEReqLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPEReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) PEReqLogic {
	return PEReqLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PEReqLogic) PEReq(req types.PrivilegeEscalationReq) error {
	payLoad := l.ctx.Value("payload").(*security.PayLoad)
	_, err := l.svcCtx.UserClient.RequestPrivilegeEscalation(l.ctx, &userclient.PrivilegeEscalationReq{
		UserId: payLoad.UserId,
		Role:   req.Role,
		Reason: req.Reason,
	})
	return err
}
