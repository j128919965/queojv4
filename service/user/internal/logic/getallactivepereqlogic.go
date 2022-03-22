package logic

import (
	"context"

	"queoj/service/user/internal/svc"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllActivePEReqLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllActivePEReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllActivePEReqLogic {
	return &GetAllActivePEReqLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllActivePEReqLogic) GetAllActivePEReq(in *user.Empty) (*user.PEList, error) {
	// todo: add your logic here and delete this line

	return &user.PEList{}, nil
}
