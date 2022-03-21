package logic

import (
	"context"

	"queoj/service/ask/ask"
	"queoj/service/ask/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type EditAskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditAskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditAskLogic {
	return &EditAskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditAskLogic) EditAsk(in *ask.AskDetail) (*ask.Empty, error) {
	// todo: add your logic here and delete this line

	return &ask.Empty{}, nil
}
