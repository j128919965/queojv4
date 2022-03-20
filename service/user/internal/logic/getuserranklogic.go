package logic

import (
	"context"

	"queoj/service/user/internal/svc"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserRankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRankLogic {
	return &GetUserRankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserRankLogic) GetUserRank(in *user.RankByUserIdReq) (*user.UserRank, error) {
	var i int32
	err := l.svcCtx.Db.Raw("select count(1) from user_accounts where point > (select point from user_accounts where id = ?)", in.Id).Scan(&i).Error
	if err != nil {
		return nil, err
	}
	return &user.UserRank{Value: i + 1 }, nil
}
