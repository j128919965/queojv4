package logic

import (
	"context"
	"queoj/service/ask/internal/model"

	"queoj/service/ask/ask"
	"queoj/service/ask/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddReplyLogic {
	return &AddReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddReplyLogic) AddReply(in *ask.ReplyDetail) (*ask.Empty, error) {
	l.svcCtx.Db.Create(ProtoReplyToModelReply(in))
	return &ask.Empty{}, nil
}

func ProtoReplyToModelReply(r *ask.ReplyDetail) * model.Reply{
	return &model.Reply{
		Id:       r.Id,
		Uid:      r.Uid,
		AskId:    r.AskId,
		Time:     r.Time,
		Nickname: r.Nickname,
		Content:  r.Content,
		IsTeacher: r.IsTeacher,
	}
}
