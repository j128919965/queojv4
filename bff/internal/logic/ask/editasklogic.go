package ask

import (
	"context"
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
