package logic

import (
	"context"
	"queoj/service/user/internal/model"

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
	var p model.PeReq
	l.svcCtx.Db.Where("user_id = ? and approval = 0" , in.Id).Find(&p)
	return &user.PrivilegeEscalationDetail{
		Id:       p.ID,
		UserId:   p.UserId,
		Role:     p.Role,
		Reason:   p.Reason,
		Approval: p.Approval,
	}, nil
}
