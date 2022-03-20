package problem

import (
	"context"
	"queoj/service/problem/problemclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetStatisticLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetStatisticLogic {
	return GetStatisticLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStatisticLogic) GetStatistic() (resp *types.AllProblemStatistic, err error) {
	statistic, err := l.svcCtx.ProblemClient.GetAllProblemStatistic(l.ctx, &problemclient.Empty{})
	if err != nil {
		return nil, err
	}
	return &types.AllProblemStatistic{
		Easy:   statistic.Easy,
		Medium: statistic.Medium,
		Hard:   statistic.Hard,
	},nil
}
