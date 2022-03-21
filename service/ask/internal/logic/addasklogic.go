package logic

import (
	"context"
	"queoj/service/ask/internal/model"

	"queoj/service/ask/ask"
	"queoj/service/ask/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddAskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAskLogic {
	return &AddAskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAskLogic) AddAsk(in *ask.AskDetail) (*ask.Empty, error) {
	l.svcCtx.Db.Create(ProtoAskToModelAsk(in))
	return &ask.Empty{}, nil
}

func ProtoAskToModelAsk(a *ask.AskDetail) *model.Ask {
	return &model.Ask{
		Id:       a.Id,
		Uid:      a.Uid,
		Time:     a.Time,
		Nickname: a.Nickname,
		Title:    a.Title,
		Content:  a.Content,
	}
}