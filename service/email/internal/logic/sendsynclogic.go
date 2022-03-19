package logic

import (
	"context"

	"queoj/service/email/email"
	"queoj/service/email/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type SendSyncLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSyncLogic {
	return &SendSyncLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSyncLogic) SendSync(in *email.SendMailReq) (*email.SendMailResp, error) {
	info := &svc.SendMailInfo{
		Address: in.Address,
		Content: in.Content,
		Subject: in.Subject,
	}
	err := l.svcCtx.SendMailSync(info)
	if err != nil {
		return nil, err
	}
	return &email.SendMailResp{
		Success: true,
	}, nil
}
