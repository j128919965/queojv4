package message

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"queoj/service/message/messageclient"

	"github.com/tal-tech/go-zero/core/logx"
	"queoj/bff/internal/svc"
)

type ReadAllLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReadAllLogic {
	return ReadAllLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadAllLogic) ReadAll() error {
	_, err := l.svcCtx.MessageClient.ReadAllMsg(l.ctx,&messageclient.MessageByUserIdReq{UserId: l.ctx.Value("payload").(*security.PayLoad).UserId})
	return err
}
