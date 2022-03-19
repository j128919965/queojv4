package user

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"queoj/service/user/user"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangePasswordLogic {
	return ChangePasswordLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req types.ChangePasswordReq) error {
	id := l.ctx.Value("payload").(*security.PayLoad).UserId
	_, err := l.svcCtx.UserClient.ChangePassword(l.ctx,&user.ChangePasswordReq{Id:id,New: req.NewPassword,Code: req.Code})
	if err != nil {
		return err
	}
	return nil
}
