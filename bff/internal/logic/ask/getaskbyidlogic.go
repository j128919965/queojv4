package ask

import (
	"context"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAskByIdLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAskByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAskByIdLogic {
	return GetAskByIdLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAskByIdLogic) GetAskById(req types.AskByIdReq) (resp *types.AskDetail, err error) {
	// todo: add your logic here and delete this line

	return
}
