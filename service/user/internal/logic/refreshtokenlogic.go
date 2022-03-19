package logic

import (
	"context"

	"queoj/service/user/internal/svc"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshTokenLogic) RefreshToken(in *user.RefreshReq) (*user.Tokens, error) {
	payload, err := l.svcCtx.TokenParser.ParseRefreshToken(in.RefreshToken)
	if err != nil {
		return nil, err
	}

	tokens, err := l.svcCtx.TokenGenerator.GenerateTokens(payload.UserId, payload.Role, payload.Refreshed+1)
	if err != nil {
		return nil, err
	}

	return &user.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
