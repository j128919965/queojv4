package user

import (
	"context"

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
	// todo: add your logic here and delete this line

	return nil
}
