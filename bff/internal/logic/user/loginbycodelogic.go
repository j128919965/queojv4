package user

import (
	"context"
	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"
	"queoj/service/user/user"
	"strconv"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginByCodeLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginByCodeLogic {
	return LoginByCodeLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByCodeLogic) LoginByCode(req types.LoginByCodeReq) (resp *types.LoginResult, err error) {
	res, err := l.svcCtx.UserClient.LoginByCode(l.ctx, &user.LoginByCodeReq{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResult{
		Info: types.UserInfo{
			Id:           strconv.FormatUint(res.Info.Id, 10),
			Nickname:     res.Info.Nickname,
			Favicon:      res.Info.Favicon,
			Phone:        res.Info.Phone,
			Email:        res.Info.Email,
			Coins:        res.Info.Coins,
			Point:        res.Info.Point,
			Introduction: res.Info.Introduction,
			Github:       res.Info.Github,
			Website:      res.Info.Website,
			Wechat:       res.Info.Wechat,
			Role:         res.Info.Role,
		},
		IsNewUser: res.IsNewUser,
		Tokens: types.Tokens{
			AccessToken:  res.Tokens.AccessToken,
			RefreshToken: res.Tokens.RefreshToken,
		},
	}, err
}
