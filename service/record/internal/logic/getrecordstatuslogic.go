package logic

import (
	"context"

	"queoj/service/record/internal/svc"
	"queoj/service/record/record"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetRecordStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecordStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecordStatusLogic {
	return &GetRecordStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRecordStatusLogic) GetRecordStatus(in *record.RecordByIdReq) (*record.RecordState, error) {
	recordById, err := l.svcCtx.GetRecordById(in.Id)
	if err != nil {
		return nil, err
	}
	return &record.RecordState{Status: recordById.Status}, nil
}
