package logic

import (
	"context"
	"queoj/service/record/internal/model"

	"queoj/service/record/internal/svc"
	"queoj/service/record/record"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetRecordByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecordByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecordByUserLogic {
	return &GetRecordByUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRecordByUserLogic) GetRecordByUser(in *record.RecordByUserReq) (*record.RecordList, error) {
	records, err := l.svcCtx.GetRecordByUserId(in.UserId)
	if err != nil {
		return nil, err
	}
	var ret []*record.RecordDetail
	for _, r := range records {
		ret = append(ret , GetRecordDetailWithOutCOde(r))
	}
	return &record.RecordList{Records: ret}, nil
}


func GetRecordDetailWithOutCOde(r *model.Record) *record.RecordDetail {
	return &record.RecordDetail{
		Id:        r.Id,
		Uid:       r.Uid,
		Time:      r.Time,
		Pid:       r.Pid,
		Language:  r.Language,
		Status:    r.Status,
		TimeUsed:  r.TimeUsed,
		SpaceUsed: r.SpaceUsed,
	}
}