package logic

import (
	"context"
	"queoj/service/user/internal/model"
	"queoj/service/user/internal/svc"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddCoinOrPointLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCoinOrPointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCoinOrPointLogic {
	return &AddCoinOrPointLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCoinOrPointLogic) AddCoinOrPoint(in *user.AddCoinOrPointReq) (*user.Result, error) {
	l.svcCtx.CoinsAndPointsSyncer.Ch <- model.UserAccount{
		ID:     in.UserId,
		Coins:  in.Coins,
		Point: in.Points,
	}
	return &user.Result{}, nil
}
