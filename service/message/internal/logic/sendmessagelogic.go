package logic

import (
	"context"
	"queoj/service/message/internal/model"
	"queoj/service/message/internal/svc"
	"queoj/service/message/message"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *message.MessageDto) (*message.Empty, error) {
	msg := model.Message{
		Receiver:  in.Receiver,
		Type:      byte(in.Type),
		Title:     in.Title,
		Content:   &in.Content,
		CreatedAt: time.Now(),
	}
	err := l.svcCtx.Db.Create(&msg).Error
	return &message.Empty{}, err
}
