package logic

import (
	"context"

	"queoj/service/ask/ask"
	"queoj/service/ask/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveAskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveAskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAskLogic {
	return &RemoveAskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveAskLogic) RemoveAsk(in *ask.AskByIdReq) (*ask.Empty, error) {
	// todo: add your logic here and delete this line

	return &ask.Empty{}, nil
}
