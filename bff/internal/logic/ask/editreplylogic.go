package ask

import (
	"context"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EditReplyLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) EditReplyLogic {
	return EditReplyLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditReplyLogic) EditReply(req types.ReplyDetail) error {
	// todo: add your logic here and delete this line

	return nil
}
