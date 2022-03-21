package ask

import (
	"context"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveReplyLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveReplyLogic {
	return RemoveReplyLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveReplyLogic) RemoveReply(req types.AskByIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
