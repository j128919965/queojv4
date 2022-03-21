package ask

import (
	"context"
	"queoj/service/ask/askclient"

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
	_, err := l.svcCtx.AskClient.RemoveReply(l.ctx, &askclient.ReplyByIdReq{Id: req.Id})
	return err
}
