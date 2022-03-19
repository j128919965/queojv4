package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/j128919965/gopkg/stringx"
	"queoj/service/user/internal/model"

	"queoj/service/user/internal/svc"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChangeUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeUserInfoLogic {
	return &ChangeUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeUserInfoLogic) ChangeUserInfo(in *user.ChangeUserInfoReq) (*user.Result, error) {
	userInfo := model.UserInfo{ID: in.UserId}
	update := model.UserInfo{}

	if in.Phone!=nil {
		if stringx.IsMobile(in.Phone) {
			update.Phone = sql.NullString{String: *in.Phone, Valid: true}
		}else {
			return nil, fmt.Errorf("手机号码 %s 错误",*in.Phone)
		}
	}

	if in.Github!=nil {
		if stringx.IsGithub(in.Github) {
			update.Github = sql.NullString{String: *in.Github, Valid: true}
		}else {
			return nil, fmt.Errorf("github %s 错误",*in.Github)
		}
	}

	if in.Website!=nil {
		if stringx.IsWebSite(in.Website) {
			update.Website = sql.NullString{String: *in.Website, Valid: true}
		}else {
			return nil, fmt.Errorf("网站格式 %s 错误",*in.Website)
		}
	}

	if in.Favicon!=nil {
		if stringx.IsWebSite(in.Favicon) {
			update.Favicon = sql.NullString{String: *in.Favicon, Valid: true}
		}else {
			return nil, fmt.Errorf("网站格式 %s 错误",*in.Website)
		}
	}

	if in.Wechat!=nil {
		if !stringx.IsBlank(in.Wechat) {
			update.Wechat = sql.NullString{String: *in.Wechat, Valid: true}
		}else {
			return nil, fmt.Errorf("微信 %s 不能为空",*in.Wechat)
		}
	}

	if in.Introduction!=nil {
		if !stringx.IsBlank(in.Introduction) {
			update.Introduction = sql.NullString{String: *in.Introduction, Valid: true}
		}else {
			return nil, fmt.Errorf("个人介绍 %s 不能为空",*in.Introduction)
		}
	}

	if in.Nickname!=nil {
		if !stringx.IsBlank(in.Nickname) {
			update.Nickname = sql.NullString{String: *in.Nickname, Valid: true}
		}else {
			return nil, fmt.Errorf("昵称 %s 不能为空",*in.Nickname)
		}
	}

	if 	err := l.svcCtx.Db.Model(&userInfo).Updates(update).Error ; err != nil{
		return nil,err
	}

	l.svcCtx.RemoveUserInfoCache(userInfo.ID)
	return &user.Result{}, nil
}
