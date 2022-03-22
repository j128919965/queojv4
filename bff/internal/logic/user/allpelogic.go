package user

import (
	"context"
	"github.com/j128919965/gopkg/errors"
	"github.com/j128919965/gopkg/security"
	"queoj/service/user/userclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllPELogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllPELogic(ctx context.Context, svcCtx *svc.ServiceContext) AllPELogic {
	return AllPELogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllPELogic) AllPE() (resp *types.PEList, err error) {
	payLoad := l.ctx.Value("payload").(*security.PayLoad)
	if payLoad.Role < 3 {
		return nil,errors.New("权限不足",400)
	}
	list, err := l.svcCtx.UserClient.GetAllActivePEReq(l.ctx, &userclient.Empty{})
	if err != nil {
		return nil, err
	}
	var pes []*types.PrivilegeEscalationDetail
	for _, p := range list.Pes {
		pes = append(pes , ProtoPeDetailToJson(p))
	}
	return &types.PEList{Pes: pes},nil
}

func ProtoPeDetailToJson(p *userclient.PrivilegeEscalationDetail) *types.PrivilegeEscalationDetail {
	return &types.PrivilegeEscalationDetail{
		Id:       p.Id,
		UserId:   p.UserId,
		Role:     p.Role,
		Reason:   p.Reason,
		Approval: p.Approval,
	}
}