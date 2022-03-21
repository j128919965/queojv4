package ask

import (
	"context"
	"queoj/service/ask/askclient"

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

func (l *EditReplyLogic) EditReply(a types.ReplyDetail) error {
	_, err := l.svcCtx.AskClient.EditReply(l.ctx, &askclient.ReplyDetail{
		Id:       a.Id,
		AskId:    a.AskId,
		Uid:      a.Uid,
		Time:     a.Time,
		Nickname: a.Nickname,
		Content:  a.Content,
	})
	return err
}
