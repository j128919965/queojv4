package user

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
