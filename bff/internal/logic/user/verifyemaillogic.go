package user

import (
	"context"
	"github.com/j128919965/gopkg/errors"
	"github.com/j128919965/gopkg/stringx"
	"queoj/service/user/user"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type VerifyEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerifyEmailLogic {
	return VerifyEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyEmailLogic) VerifyEmail(req types.VerifyEmailReq) error {
	email := req.Email
	if !stringx.IsEmail(&email) {
		return errors.New("email 格式不正确！",400)
	}

	_, err := l.svcCtx.UserClient.SendVerifyEmail(l.ctx,&user.UserInfoReqByEmail{Email: req.Email})
	if err != nil {
		return err
	}

	return nil
}
