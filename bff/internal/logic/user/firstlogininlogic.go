package user

import (
	"context"
	"fmt"
	"github.com/j128919965/gopkg/security"
	"queoj/service/user/user"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
	"queoj/bff/internal/svc"
)

type FirstLoginInLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFirstLoginInLogic(ctx context.Context, svcCtx *svc.ServiceContext) FirstLoginInLogic {
	return FirstLoginInLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FirstLoginInLogic) FirstLoginIn() error {
	userId := l.ctx.Value("payload").(*security.PayLoad).UserId
	key := fmt.Sprintf("first:%d:%d", time.Now().Day(), userId)
	exists, err := l.svcCtx.Redis.Exists(key)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	_, err = l.svcCtx.UserClient.AddCoinOrPoint(l.ctx, &user.AddCoinOrPointReq{
		UserId: userId,
		Coins:  3,
	})
	if err != nil {
		return err
	}

	_, err = l.svcCtx.Redis.SetnxEx(key, "1", 60*60*24)
	if err != nil {
		return err
	}

	return nil
}
