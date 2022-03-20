package user

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"queoj/service/user/userclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserRankLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserRankLogic {
	return GetUserRankLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRankLogic) GetUserRank() (resp *types.UserRank, err error) {
	rank, err := l.svcCtx.UserClient.GetUserRank(l.ctx,&userclient.RankByUserIdReq{Id: l.ctx.Value("payload").(*security.PayLoad).UserId})
	if err != nil {
		return nil, err
	}
	return &types.UserRank{Rank: rank.Value},nil
}
