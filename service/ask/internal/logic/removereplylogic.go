package logic

import (
	"context"

	"queoj/service/ask/ask"
	"queoj/service/ask/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveReplyLogic {
	return &RemoveReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveReplyLogic) RemoveReply(in *ask.ReplyByIdReq) (*ask.Empty, error) {
	// todo: add your logic here and delete this line

	return &ask.Empty{}, nil
}
