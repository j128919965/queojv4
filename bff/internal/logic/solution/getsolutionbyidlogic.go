package solution

import (
	"context"
	"queoj/service/solution/solutionclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetSolutionByIdLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSolutionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetSolutionByIdLogic {
	return GetSolutionByIdLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSolutionByIdLogic) GetSolutionById(req types.SolutionByIdReq) (resp *types.SolutionDetail, err error) {
	d, err := l.svcCtx.SolutionClient.GetSolutionDetail(l.ctx, &solutionclient.SolutionByIdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.SolutionDetail{
		Id:       d.Id,
		Uid:      d.Uid,
		Pid:      d.Pid,
		Time:     d.Time,
		Nickname: d.Nickname,
		Title:    d.Title,
		Summary:  d.Summary,
		Content:  d.Content,
	},nil
}
