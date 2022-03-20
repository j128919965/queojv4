package logic

import (
	"context"

	"queoj/service/problem/internal/svc"
	"queoj/service/problem/problem"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllProblemStatisticLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllProblemStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllProblemStatisticLogic {
	return &GetAllProblemStatisticLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllProblemStatisticLogic) GetAllProblemStatistic(in *problem.Empty) (*problem.AllProblemStatistic, error) {
	var easy,medium,hard int32
	l.svcCtx.Db.Raw("select count(1) from problems where level = 1").Scan(&easy)
	l.svcCtx.Db.Raw("select count(1) from problems where level = 2").Scan(&medium)
	l.svcCtx.Db.Raw("select count(1) from problems where level = 3").Scan(&hard)
	return &problem.AllProblemStatistic{
		Easy:   easy,
		Medium: medium,
		Hard:   hard,
	}, nil
}
