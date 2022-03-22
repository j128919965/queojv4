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
	payLoad := l.ctx.Value("payload").(*security.PayLoad)
	if payLoad.Role < 2 {
		return errors.New("权限不足",400)
	}
	_, err := l.svcCtx.AskClient.EditReply(l.ctx, &askclient.ReplyDetail{
		Id:       a.Id,
		AskId:    a.AskId,
		Uid:      a.Uid,
		Time:     a.Time,
		Nickname: a.Nickname,
		Content:  a.Content,
		IsTeacher: a.IsTeacher,
	})
	return err
}
