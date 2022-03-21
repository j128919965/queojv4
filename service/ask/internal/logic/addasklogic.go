package logic

import (
	"context"

	"queoj/service/ask/ask"
	"queoj/service/ask/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddAskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAskLogic {
	return &AddAskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAskLogic) AddAsk(in *ask.AskDetail) (*ask.Empty, error) {
	// todo: add your logic here and delete this line

	return &ask.Empty{}, nil
}
