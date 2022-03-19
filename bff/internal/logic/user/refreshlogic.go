package user

import (
	"context"
	"queoj/service/user/user"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RefreshLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) RefreshLogic {
	return RefreshLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh(req types.RefreshReq) (resp *types.Tokens, err error) {
	tokens, err := l.svcCtx.UserClient.RefreshToken(l.ctx, &user.RefreshReq{RefreshToken: req.RefreshToken})
	if err != nil {
		return nil, err
	}
	return &types.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	},nil
}
