package logic

import (
	"context"

	"queoj/service/message/internal/svc"
	"queoj/service/message/message"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReadMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReadMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadMsgLogic {
	return &ReadMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReadMsgLogic) ReadMsg(in *message.MessageByIdReq) (*message.Empty, error) {
	return &message.Empty{}, l.svcCtx.ReadMsg(in.Id)
}
