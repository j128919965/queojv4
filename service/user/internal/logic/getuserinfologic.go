package logic

import (
	"context"
	"github.com/j128919965/gopkg/errors"
	"github.com/j128919965/gopkg/stringx"
	"queoj/service/user/internal/svc"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.UserInfoReq) (*user.UserInfo, error) {
	if in.UserId==nil && in.Email == nil {
		return nil,errors.New("id 或 email 至少有一个不能为空！",400)
	}

	if in.UserId == nil{
		if !stringx.IsEmail(in.Email){
			return nil,errors.New("email格式不正确",400)
		}
		id ,err := l.svcCtx.GetUidByEmail(*in.Email)
		if err != nil {
			return nil, err
		}
		in.UserId = &id
	}

	info ,err := l.svcCtx.GetUserInfo(*in.UserId)
	if err != nil {
		return nil,err
	}

	account ,err := l.svcCtx.GetUserAccount(*in.UserId)
	if err != nil {
		return nil,err
	}

	return &user.UserInfo{
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
	}, nil
}
