package ask

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"queoj/service/ask/askclient"
	"time"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddAskLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAskLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddAskLogic {
	return AddAskLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAskLogic) AddAsk(req types.AskAddReq) error {
	payLoad := l.ctx.Value("payload").(*security.PayLoad)
	userId := payLoad.UserId
	_, err := l.svcCtx.AskClient.AddAsk(l.ctx,&askclient.AskDetail{
		Uid:      userId,
		Time:     time.Now().Unix(),
		Nickname: req.Nickname,
		Title:    req.Title,
		Content:  req.Content,
	})
	return err
}
