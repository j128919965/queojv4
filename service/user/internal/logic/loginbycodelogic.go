package logic

import (
	"context"
	"database/sql"
	"github.com/j128919965/gopkg/errors"
	"github.com/j128919965/gopkg/stringx"
	"gorm.io/gorm"
	"queoj/service/message/messageclient"
	"queoj/service/user/internal/model"
	"time"

	"queoj/service/user/internal/svc"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginByCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByCodeLogic {
	return &LoginByCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginByCodeLogic) LoginByCode(in *user.LoginByCodeReq) (*user.LoginResult, error) {
	if !stringx.IsEmail(&in.Email) {
		return nil, errors.New("邮箱格式不正确", 400)
	}

	code, err := l.svcCtx.Redis.Get(codeEmailPrefix + in.Email)
	if err != nil {
		return nil, err
	}

	if code == "" || code != in.Code {
		return nil, errors.New("验证码无效或已过期", 400)
	}

	u := model.User{}
	err = l.svcCtx.Db.Where(&model.User{Email: in.Email}).First(&u).Error

	resp := &user.LoginResult{
		Info:      nil,
		Tokens:    nil,
		IsNewUser: false,
	}

	if err != nil && err == gorm.ErrRecordNotFound {
		// new User
		u.Email = in.Email
		l.svcCtx.Db.Create(&u)
		info := model.UserInfo{
			ID:       u.ID,
			Email:    u.Email,
			Nickname: sql.NullString{String: stringx.RandomNickName(), Valid: true},
		}
		l.svcCtx.Db.Create(&info)
		account := model.UserAccount{
			ID: u.ID,
		}
		l.svcCtx.Db.Create(&account)

		resp.IsNewUser = true
		resp.Info = &user.UserInfo{Email: u.Email, Nickname: &info.Nickname.String}

		_, err := l.svcCtx.MessageClient.SendMessage(l.ctx, &messageclient.MessageDto{
			Receiver: u.ID,
			Time:     time.Now().Unix(),
			Type:     0,
			Title:    "欢迎来到QueOj!",
			Content:  "欢迎来到QueOj!\n这里有丰富多样的题型，舒适的编码体验，\n每道题还有题解可以阅读。\n\n享受编程的乐趣吧！",
		})
		if err != nil {
			l.Error(err)
		}
	} else {
		info, err := l.svcCtx.GetUserInfo(u.ID)
		if err != nil {
			return nil, err
		}

		account, err := l.svcCtx.GetUserAccount(u.ID)
		if err != nil {
			return nil, err
		}

		resp.Info = &user.UserInfo{
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
			Role:         int32(u.Role),
		}
	}

	tokens, err := l.svcCtx.TokenGenerator.GenerateTokens(u.ID, int32(u.Role), 0)
	if err != nil {
		return nil, err
	}

	resp.Tokens = &user.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	return resp, nil
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}