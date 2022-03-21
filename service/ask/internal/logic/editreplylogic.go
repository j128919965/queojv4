package logic

import (
	"context"

	"queoj/service/ask/ask"
	"queoj/service/ask/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type EditReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditReplyLogic {
	return &EditReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditReplyLogic) EditReply(in *ask.ReplyDetail) (*ask.Empty, error) {
	l.svcCtx.Db.Save(ProtoReplyToModelReply(in))
	return &ask.Empty{}, nil
}
