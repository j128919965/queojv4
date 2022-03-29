package logic

import (
	"context"
	"fmt"
	"queoj/service/record/internal/svc"
	"queoj/service/record/record"
	"strconv"

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
	get, err := l.svcCtx.Redis.Get(fmt.Sprintf("r:s:%d", in.GetId()))
	if err == nil {
		if get != "" {
			parseUint, err := strconv.ParseUint(get, 10, 32)
			if err == nil {
				return &record.RecordState{Status: uint32(parseUint)} , nil
			}
		}
	}

	recordById, err := l.svcCtx.GetRecordById(in.Id)
	if err != nil {
		return nil, err
	}
	return &record.RecordState{Status: recordById.Status}, nil
}
