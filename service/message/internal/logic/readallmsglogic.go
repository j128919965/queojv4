package logic

import (
	"context"

	"queoj/service/message/internal/svc"
	"queoj/service/message/message"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReadAllMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReadAllMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadAllMsgLogic {
	return &ReadAllMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReadAllMsgLogic) ReadAllMsg(in *message.MessageByUserIdReq) (*message.Empty, error) {
	return &message.Empty{}, l.svcCtx.ReadAllMsg(in.UserId)
}
