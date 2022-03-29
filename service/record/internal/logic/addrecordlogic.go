package logic

import (
	"context"
	"queoj/service/record/internal/model"

	"queoj/service/record/internal/svc"
	"queoj/service/record/record"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRecordLogic {
	return &AddRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRecordLogic) AddRecord(in *record.RecordDetail) (*record.RecordByIdReq, error) {
	r := &model.Record{
		Uid:       in.Uid,
		Pid:       in.Pid,
		Time:      in.Time,
		Language:  in.Language,
		Status:    0,
		Code:      in.Code,
	}
	err := l.svcCtx.InsertRecord(r)
	if err != nil {
		return nil, err
	}

	// 异步判题
	l.svcCtx.SubmitRecord(r)

	return &record.RecordByIdReq{
		Id: r.Id,
	}, nil
}
