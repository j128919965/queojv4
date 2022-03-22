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

func (l *EditAskLogic) EditAsk(a types.AskDetail) error {
	payLoad := l.ctx.Value("payload").(*security.PayLoad)
	if payLoad.Role < 2 {
		return errors.New("权限不足",400)
	}
	_, err := l.svcCtx.AskClient.EditAsk(l.ctx, &askclient.AskDetail{
		Id:       a.Id,
		Uid:      a.Uid,
		Time:     a.Time,
		Nickname: a.Nickname,
		Title:    a.Title,
		Content:  a.Content,
	})
	return err
}
