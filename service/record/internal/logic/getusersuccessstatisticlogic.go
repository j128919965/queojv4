package logic

import (
	"context"
	"queoj/service/record/internal/model"

	"queoj/service/record/internal/svc"
	"queoj/service/record/record"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserSuccessStatisticLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserSuccessStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSuccessStatisticLogic {
	return &GetUserSuccessStatisticLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserSuccessStatisticLogic) GetUserSuccessStatistic(in *record.RecordByUserReq) (*record.SuccessStatistic, error) {
	var s model.UserSuccessStatistic
	err := l.svcCtx.Db.Where("uid = ?", in.UserId).Find(&s).Error
	if err != nil {
		return nil, err
	}
	return &record.SuccessStatistic{
		Easy:   s.Easy,
		Medium: s.Medium,
		Hard:   s.Hard,
	}, nil
}
