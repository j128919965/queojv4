package solution

import (
	"context"
	"queoj/service/solution/solutionclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllSolutionByPidLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllSolutionByPidLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllSolutionByPidLogic {
	return GetAllSolutionByPidLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllSolutionByPidLogic) GetAllSolutionByPid(req types.AllSolutionByPidReq) (resp *types.SolutionList, err error) {
	ss, err := l.svcCtx.SolutionClient.GetAllSolutionByPid(l.ctx, &solutionclient.AllSolutionByPidReq{Pid: req.Pid})
	if err != nil {
		return nil, err
	}

	var ret []*types.SolutionSummary

	for _,s := range ss.GetSolutions() {
		ret = append(ret , &types.SolutionSummary{
			Id:       s.Id,
			Time:     s.Time,
			Nickname: s.Nickname,
			Title:    s.Title,
			Summary:  s.Summary,
		})
	}

	return &types.SolutionList{Solutions: ret},nil
}
