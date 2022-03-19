package user

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChangeUserInfoLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangeUserInfoLogic {
	return ChangeUserInfoLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeUserInfoLogic) ChangeUserInfo(req types.ChangeUserInfoReq) error {
	rpcReq := &user.ChangeUserInfoReq{
		UserId:       l.ctx.Value("payload").(*security.PayLoad).UserId,
		Nickname:     req.Nickname,
		Favicon:      req.Favicon,
		Phone:        req.Phone,
		Introduction: req.Introduction,
		Github:       req.Github,
		Website:      req.Website,
		Wechat:       req.Wechat,
	}

	_, err := l.svcCtx.UserClient.ChangeUserInfo(l.ctx, rpcReq)
	if err != nil {
		return err
	}
	return nil
}
