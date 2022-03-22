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
	payLoad := l.ctx.Value("payload").(*security.PayLoad)
	if payLoad.Role < 2 {
		return errors.New("权限不足",400)
	}
	_, err := l.svcCtx.AskClient.RemoveAsk(l.ctx, &askclient.AskByIdReq{Id: req.Id})
	return err
}
