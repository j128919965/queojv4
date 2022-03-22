package logic

import (
	"context"

	"queoj/service/user/internal/svc"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type ApprovalPEReqLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApprovalPEReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApprovalPEReqLogic {
	return &ApprovalPEReqLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ApprovalPEReqLogic) ApprovalPEReq(in *user.ApprovalPrivilegeEscalationReq) (*user.Result, error) {
	// todo: add your logic here and delete this line

	return &user.Result{}, nil
}
