package logic

import (
	"context"
	"queoj/service/record/internal/model"

	"queoj/service/record/internal/svc"
	"queoj/service/record/record"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetRecordByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecordByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecordByIdLogic {
	return &GetRecordByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRecordByIdLogic) GetRecordById(in *record.RecordByIdReq) (*record.RecordDetail, error) {
	recordById, err := l.svcCtx.GetRecordById(in.Id)
	if err != nil {
		return nil, err
	}
	return GetRecordDetail(recordById), nil
}

func GetRecordDetail(r *model.Record) *record.RecordDetail {
	return &record.RecordDetail{
		Id:        r.Id,
		Uid:       r.Uid,
		Time:      r.Time,
		Pid:       r.Pid,
		Language:  r.Language,
		Status:    r.Status,
		TimeUsed:  r.TimeUsed,
		SpaceUsed: r.SpaceUsed,
		Code:      r.Code,
	}
}