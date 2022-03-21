package ask

import (
	"context"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddAskLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAskLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddAskLogic {
	return AddAskLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAskLogic) AddAsk(req types.AskAddReq) error {
	// todo: add your logic here and delete this line

	return nil
}
