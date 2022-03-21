package ask

import (
	"context"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveAskLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveAskLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveAskLogic {
	return RemoveAskLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveAskLogic) RemoveAsk(req types.AskByIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
