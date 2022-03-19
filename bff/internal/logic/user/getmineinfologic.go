package user

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"queoj/service/user/user"
	"strconv"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetMineInfoLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMineInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMineInfoLogic {
	return GetMineInfoLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMineInfoLogic) GetMineInfo() (resp *types.UserInfo, err error) {
	id := l.ctx.Value("payload").(*security.PayLoad).UserId
	info, err := l.svcCtx.UserClient.GetUserInfo(l.ctx, &user.UserInfoReq{UserId: &id})
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
