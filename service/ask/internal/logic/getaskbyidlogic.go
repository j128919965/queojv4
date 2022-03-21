package logic

import (
	"context"

	"queoj/service/ask/ask"
	"queoj/service/ask/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAskByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAskByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAskByIdLogic {
	return &GetAskByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAskByIdLogic) GetAskById(in *ask.AskByIdReq) (*ask.AskDetail, error) {
	id, err := l.svcCtx.GetAskById(in.GetId())
	if err != nil {
		return nil, err
	}
	return ModelAskToProtoAsk(id), nil
}
