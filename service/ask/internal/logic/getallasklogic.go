package logic

import (
	"context"
	"queoj/service/ask/internal/model"

	"queoj/service/ask/ask"
	"queoj/service/ask/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllAskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllAskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllAskLogic {
	return &GetAllAskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllAskLogic) GetAllAsk(_ *ask.Empty) (*ask.AskList, error) {
	allAsk := l.svcCtx.GetAllAsk()
	var res []*ask.AskSummary
	for _, a := range allAsk {
		res = append(res,ModelAskToProtoAskSummary(a))
	}
	return &ask.AskList{Asks: res}, nil
}

func ModelAskToProtoAsk(a *model.Ask) *ask.AskDetail{
	return &ask.AskDetail{
		Id:       a.Id,
		Uid:      a.Uid,
		Time:     a.Time,
		Nickname: a.Nickname,
		Title:    a.Title,
		Content:  a.Content,
	}
}

func ModelAskToProtoAskSummary(a *model.Ask) *ask.AskSummary{
	return &ask.AskSummary{
		Id:       a.Id,
		Uid:      a.Uid,
		Time:     a.Time,
		Nickname: a.Nickname,
		Title:    a.Title,
	}
}

