package user

import (
	"context"
	"fmt"
	"queoj/service/user/user"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginByPasswordLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginByPasswordLogic {
	return LoginByPasswordLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByPasswordLogic) LoginByPassword(req types.LoginByPasswordReq) (resp *types.LoginResult, err error) {
	result, err := l.svcCtx.UserClient.LoginByPassword(l.ctx, &user.LoginByPasswordReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.LoginResult{
		Info: types.UserInfo{
			Id:           fmt.Sprintf("%d", result.Info.Id),
			Nickname:     result.Info.Nickname,
			Favicon:      result.Info.Favicon,
			Phone:        result.Info.Phone,
			Email:        result.Info.Email,
			Coins:        result.Info.Coins,
			Point:        result.Info.Point,
			Introduction: result.Info.Introduction,
			Github:       result.Info.Github,
			Website:      result.Info.Website,
			Wechat:       result.Info.Wechat,
			Role:         result.Info.Role,
		},
		IsNewUser: result.IsNewUser,
		Tokens: types.Tokens{
			AccessToken:  result.Tokens.AccessToken,
			RefreshToken: result.Tokens.RefreshToken,
		},
	}
	return resp, nil
}
