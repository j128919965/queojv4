package logic

import (
	"context"

	"queoj/service/user/internal/svc"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserPrivilegeEscalationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPrivilegeEscalationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPrivilegeEscalationLogic {
	return &GetUserPrivilegeEscalationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserPrivilegeEscalationLogic) GetUserPrivilegeEscalation(in *user.RankByUserIdReq) (*user.PrivilegeEscalationDetail, error) {
	// todo: add your logic here and delete this line

	return &user.PrivilegeEscalationDetail{}, nil
}
