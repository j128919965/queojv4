package record

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"queoj/service/record/recordclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GeySuccessStatisticByUserLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGeySuccessStatisticByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GeySuccessStatisticByUserLogic {
	return GeySuccessStatisticByUserLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GeySuccessStatisticByUserLogic) GeySuccessStatisticByUser() (resp *types.SuccessStatistic, err error) {
	s, err := l.svcCtx.RecordClient.GetUserSuccessStatistic(l.ctx,&recordclient.RecordByUserReq{UserId: l.ctx.Value("payload").(*security.PayLoad).UserId})
	if err != nil {
		return nil, err
	}
	return &types.SuccessStatistic{
		Easy:   s.Easy,
		Medium: s.Medium,
		Hard:   s.Hard,
	},nil
}
