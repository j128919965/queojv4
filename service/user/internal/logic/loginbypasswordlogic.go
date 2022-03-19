package logic

import (
	"context"
	"github.com/j128919965/gopkg/errors"
	"github.com/j128919965/gopkg/stringx"
	"queoj/service/user/internal/model"

	"queoj/service/user/internal/svc"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginByPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPasswordLogic {
	return &LoginByPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginByPasswordLogic) LoginByPassword(in *user.LoginByPasswordReq) (*user.LoginResult, error) {
	u := model.User{}
	err := l.svcCtx.Db.Where(&model.User{Email: in.Email}).First(&u).Error
	if err != nil {
		return nil, err
	}

	if !u.Password.Valid{
		return nil, errors.New("该用户暂未设置密码，不能使用密码登录",400)
	}

	pass := stringx.Encrypt(l.svcCtx.Config.Salt, in.Password)
	if pass != u.Password.String {
		return nil, errors.New("邮箱或密码错误",400)
	}

	resp := &user.LoginResult{
		Info:      nil,
		Tokens:    nil,
		IsNewUser: false,
	}

	tokens, err := l.svcCtx.TokenGenerator.GenerateTokens(u.ID, int32(u.Role), 0)
	if err != nil {
		return nil, err
	}

	id := u.ID

	info ,err :=l.svcCtx.GetUserInfo(id)
	if err != nil {
		return nil,err
	}

	account ,err := l.svcCtx.GetUserAccount(id)
	if err != nil {
		return nil,err
	}

	resp.Info =  &user.UserInfo{
		Id:           info.ID,
		Nickname:     stringx.NullStringToPtr(info.Nickname),
		Favicon:      stringx.NullStringToPtr(info.Favicon),
		Phone:        stringx.NullStringToPtr(info.Phone),
		Email:        info.Email,
		Coins:        account.Coins,
		Point:        account.Point,
		Introduction: stringx.NullStringToPtr(info.Introduction),
		Github:       stringx.NullStringToPtr(info.Github),
		Website:      stringx.NullStringToPtr(info.Website),
		Wechat:       stringx.NullStringToPtr(info.Wechat),
	}

	resp.Tokens = &user.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	return resp, nil
}
