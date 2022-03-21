package logic

import (
	"context"

	"queoj/service/ask/ask"
	"queoj/service/ask/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllAskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllAskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllAskLogic {
	return &GetAllAskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllAskLogic) GetAllAsk(in *ask.Empty) (*ask.AskList, error) {
	// todo: add your logic here and delete this line

	return &ask.AskList{}, nil
}
