package logic

import (
	"context"

	"queoj/service/email/email"
	"queoj/service/email/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type SendAsyncLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendAsyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendAsyncLogic {
	return &SendAsyncLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendAsyncLogic) SendAsync(in *email.SendMailReq) (*email.SendMailResp, error) {
	info := &svc.SendMailInfo{
		Address: in.Address,
		Content: in.Content,
		Subject: in.Subject,
	}
	err := l.svcCtx.SendMailAsync(info)
	if err != nil {
		return nil, err
	}
	return &email.SendMailResp{
		Success: true,
	}, nil
}
