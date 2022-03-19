package logic

import (
	"context"

	"queoj/service/message/internal/svc"
	"queoj/service/message/message"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetMessagePageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessagePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessagePageLogic {
	return &GetMessagePageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessagePageLogic) GetMessagePage(in *message.MessageReqByReceiver) (*message.MessagePage, error) {
	var pageFrom int32
	if in.Page.PageNumber != nil {
		pageFrom = (*in.Page.PageNumber) * in.Page.PageSize
	}

	msgs, err := l.svcCtx.GetNoReadMessagesByUser(in.Receiver, in.GetNoread(), int(in.Page.PageSize), &pageFrom, in.Page.LastId)
	if err != nil {
		return nil, err
	}

	page := &message.MessagePage{}

	if in.Page.HasCount {
		cnt,err := l.svcCtx.GetTotalCount(in.Receiver,in.GetNoread())
		if err != nil {
			return nil, err
		}
		page.TotalCount = int32(cnt)
	}

	for _, msg := range msgs {
		page.Messages = append(page.Messages , &message.MessageDto{
			Id:        msg.Id,
			Receiver:  msg.Receiver,
			Time:      msg.CreatedAt.Unix(),
			Read:      msg.Read,
			Type:      int32(msg.Type),
			Title:     msg.Title,
		})
	}


	return page, nil
}
