package ask

import (
	"context"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddReplyLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddReplyLogic {
	return AddReplyLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddReplyLogic) AddReply(req types.ReplyAddReq) error {
	// todo: add your logic here and delete this line

	return nil
}
