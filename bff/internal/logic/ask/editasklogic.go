package ask

import (
	"context"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EditAskLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditAskLogic(ctx context.Context, svcCtx *svc.ServiceContext) EditAskLogic {
	return EditAskLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditAskLogic) EditAsk(req types.AskDetail) error {
	// todo: add your logic here and delete this line

	return nil
}
