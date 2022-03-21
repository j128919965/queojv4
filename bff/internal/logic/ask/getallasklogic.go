package ask

import (
	"context"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllAskLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllAskLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllAskLogic {
	return GetAllAskLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllAskLogic) GetAllAsk() (resp *types.AskList, err error) {
	// todo: add your logic here and delete this line

	return
}
