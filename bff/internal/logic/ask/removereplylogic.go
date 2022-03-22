package ask

import (
	"context"
	"github.com/j128919965/gopkg/errors"
	"github.com/j128919965/gopkg/security"
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
	payLoad := l.ctx.Value("payload").(*security.PayLoad)
	if payLoad.Role < 2 {
		return errors.New("权限不足",400)
	}
	_, err := l.svcCtx.AskClient.RemoveReply(l.ctx, &askclient.ReplyByIdReq{Id: req.Id})
	return err
}
