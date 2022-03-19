package message

import (
	"context"
	"queoj/service/message/messageclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReadMsgLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReadMsgLogic {
	return ReadMsgLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadMsgLogic) ReadMsg(req types.MessageByIdReq) error {
	_, err := l.svcCtx.MessageClient.ReadMsg(l.ctx,&messageclient.MessageByIdReq{Id: req.Id})
	return err
}
