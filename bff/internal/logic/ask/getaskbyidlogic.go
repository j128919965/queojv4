package ask

import (
	"context"
	"queoj/service/ask/askclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAskByIdLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAskByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAskByIdLogic {
	return GetAskByIdLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAskByIdLogic) GetAskById(req types.AskByIdReq) (resp *types.AskDetail, err error) {
	a, err := l.svcCtx.AskClient.GetAskById(l.ctx, &askclient.AskByIdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.AskDetail{
		Id:       a.Id,
		Uid:      a.Uid,
		Time:     a.Time,
		Nickname: a.Nickname,
		Title:    a.Title,
		Content:  a.Content,
	},nil
}
