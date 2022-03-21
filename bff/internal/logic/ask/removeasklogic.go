package ask

import (
	"context"
	"queoj/service/ask/askclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveAskLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveAskLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveAskLogic {
	return RemoveAskLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveAskLogic) RemoveAsk(req types.AskByIdReq) error {
	_, err := l.svcCtx.AskClient.RemoveAsk(l.ctx, &askclient.AskByIdReq{Id: req.Id})
	return err
}
