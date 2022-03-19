package user

import (
	"context"
	"queoj/service/user/user"
	"strconv"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserInfoLogic {
	return GetUserInfoLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req types.UserInfoReq) (resp *types.UserInfo, err error) {
	info, err := l.svcCtx.UserClient.GetUserInfo(l.ctx, &user.UserInfoReq{Email: &req.Email})
	if err != nil {
		return nil, err
	}
	return &types.UserInfo{
		Id:           strconv.FormatUint(info.Id,10),
		Nickname:     info.Nickname,
		Favicon:      info.Favicon,
		Phone:        info.Phone,
		Email:        info.Email,
		Coins:        info.Coins,
		Point:        info.Point,
		Introduction: info.Introduction,
		Github:       info.Github,
		Website:      info.Website,
		Wechat:       info.Wechat,
	},nil
}
