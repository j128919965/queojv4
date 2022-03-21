package ask

import (
	"context"
	"queoj/service/ask/askclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllAskLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllAskLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllAskLogic {
	return GetAllAskLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllAskLogic) GetAllAsk() (resp *types.AskList, err error) {
	as, err := l.svcCtx.AskClient.GetAllAsk(l.ctx, &askclient.Empty{})
	if err != nil {
		return nil, err
	}

	var ret []*types.AskSummary
	for _, a := range as.Asks {
		ret = append(ret , &types.AskSummary{
			Id:       a.Id,
			Uid:      a.Uid,
			Time:     a.Time,
			Nickname: a.Nickname,
			Title:    a.Title,
		})
	}

	return &types.AskList{Asks: ret}, nil
}
