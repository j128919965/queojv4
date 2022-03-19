package logic

import (
	"context"

	"queoj/service/message/internal/svc"
	"queoj/service/message/message"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetMessageByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageByIdLogic {
	return &GetMessageByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageByIdLogic) GetMessageById(in *message.MessageByIdReq) (*message.MessageDto, error) {
	msg, err := l.svcCtx.GetMessageById(in.Id)
	if err != nil {
		l.Error(err)
		return nil,err
	}
	return &message.MessageDto{
		Id:        msg.Id,
		Receiver:  msg.Receiver,
		Time:      msg.CreatedAt.Unix(),
		Read:      msg.Read,
		Type:      int32(msg.Type),
		Title:     msg.Title,
		Content:   *msg.Content,
	},nil
}
