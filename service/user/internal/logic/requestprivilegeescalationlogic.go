package logic

import (
	"context"
	"queoj/service/user/internal/model"

	"queoj/service/user/internal/svc"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type RequestPrivilegeEscalationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRequestPrivilegeEscalationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RequestPrivilegeEscalationLogic {
	return &RequestPrivilegeEscalationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RequestPrivilegeEscalationLogic) RequestPrivilegeEscalation(in *user.PrivilegeEscalationReq) (*user.Result, error) {
	l.svcCtx.Db.Create(&model.PeReq{
		ID:       0,
		UserId:   in.UserId,
		Role:     in.Role,
		Reason:   in.Reason,
		Approval: 0,
	})
	return &user.Result{}, nil
}
