package logic

import (
	"context"
	"github.com/j128919965/gopkg/errors"
	"github.com/j128919965/gopkg/stringx"
	"queoj/service/email/email"

	"queoj/service/user/internal/svc"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type SendVerifyEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendVerifyEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVerifyEmailLogic {
	return &SendVerifyEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const sendMailPrefix = "em_send:"

const sendMailExpire = 59

const codeEmailPrefix = "em_code:"

const codeExpire = 5 * 60

func (l *SendVerifyEmailLogic) SendVerifyEmail(in *user.UserInfoReqByEmail) (*user.Result, error) {
	code := stringx.GenerateNumCode(4)

	what, err := l.svcCtx.Redis.Get(sendMailPrefix + in.Email)
	if err == nil && what!=""{
		return nil, errors.New("访问过于频繁，请稍后重试", 400)
	}
	if err != nil {
		return nil, err
	}

	l.svcCtx.Redis.Setex(sendMailPrefix+in.Email, "1", sendMailExpire)
	l.svcCtx.Redis.Setex(codeEmailPrefix+in.Email, code, codeExpire)

	_, err = l.svcCtx.EmailClient.SendAsync(l.ctx, &email.SendMailReq{
		Address: in.Email,
		Content: "您好，您的验证码为：" + code + " 。 请在五分钟内使用。",
		Subject: "邮箱验证",
	})
	if err != nil {
		return nil, err
	}

	return &user.Result{}, nil
}