package user

import (
	"context"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllPELogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllPELogic(ctx context.Context, svcCtx *svc.ServiceContext) AllPELogic {
	return AllPELogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllPELogic) AllPE() (resp *types.PEList, err error) {
	// todo: add your logic here and delete this line

	return
}
