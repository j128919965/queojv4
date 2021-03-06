// Code generated by goctl. DO NOT EDIT!
// Source: email.proto

package emailclient

import (
	"context"

	"queoj/service/email/email"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	SendMailReq  = email.SendMailReq
	SendMailResp = email.SendMailResp

	Email interface {
		SendAsync(ctx context.Context, in *SendMailReq) (*SendMailResp, error)
		SendSync(ctx context.Context, in *SendMailReq) (*SendMailResp, error)
	}

	defaultEmail struct {
		cli zrpc.Client
	}
)

func NewEmail(cli zrpc.Client) Email {
	return &defaultEmail{
		cli: cli,
	}
}

func (m *defaultEmail) SendAsync(ctx context.Context, in *SendMailReq) (*SendMailResp, error) {
	client := email.NewEmailClient(m.cli.Conn())
	return client.SendAsync(ctx, in)
}

func (m *defaultEmail) SendSync(ctx context.Context, in *SendMailReq) (*SendMailResp, error) {
	client := email.NewEmailClient(m.cli.Conn())
	return client.SendSync(ctx, in)
}
