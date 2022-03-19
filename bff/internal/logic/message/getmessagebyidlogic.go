package message

import (
	"context"
	"queoj/service/message/messageclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetMessageByIdLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMessageByIdLogic {
	return GetMessageByIdLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageByIdLogic) GetMessageById(req types.MessageByIdReq) (resp *types.Message, err error) {
	m, err := l.svcCtx.MessageClient.GetMessageById(l.ctx, &messageclient.MessageByIdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	msg := &types.Message{
		Id:        m.Id,
		Receiver:  m.Receiver,
		Time:      m.Time,
		Read:      m.Read,
		Type: byte(m.Type),
		Title:     m.Title,
		Content:   m.Content,
	}
	return msg,nil
}
