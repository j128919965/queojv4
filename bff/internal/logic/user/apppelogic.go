package user

import (
	"context"

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
	// todo: add your logic here and delete this line

	return nil
}
