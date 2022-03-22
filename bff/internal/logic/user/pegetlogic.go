package user

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"queoj/service/user/userclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PEGetLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPEGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) PEGetLogic {
	return PEGetLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PEGetLogic) PEGet() (resp *types.PrivilegeEscalationDetail, err error) {
	payLoad := l.ctx.Value("payload").(*security.PayLoad)
	escalation, err := l.svcCtx.UserClient.GetUserPrivilegeEscalation(l.ctx, &userclient.RankByUserIdReq{Id: payLoad.UserId})
	if err != nil {
		return nil, err
	}
	return ProtoPeDetailToJson(escalation),nil
}

