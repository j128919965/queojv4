package logic

import (
	"context"
	"queoj/service/user/internal/model"

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
	var ps []*model.PeReq
	l.svcCtx.Db.Where("approval = 0").Find(&ps)
	var ret []*user.PrivilegeEscalationDetail
	for _, p := range ps {
		ret = append(ret , &user.PrivilegeEscalationDetail{
			Id:       p.ID,
			UserId:   p.UserId,
			Role:     p.Role,
			Reason:   p.Reason,
			Approval: p.Approval,
		})
	}
	return &user.PEList{Pes: ret}, nil
}
