package record

import (
	"context"
	"queoj/service/record/recordclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetByIdLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetByIdLogic {
	return GetByIdLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByIdLogic) GetById(req types.RecordByIdReq) (resp *types.RecordDetail, err error) {
	id, err := l.svcCtx.RecordClient.GetRecordById(l.ctx, &recordclient.RecordByIdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return ProroToDetail(id), nil
}

func ProroToDetail(d *recordclient.RecordDetail) *types.RecordDetail {
	return &types.RecordDetail{
		Id:        d.Id,
		Uid:       d.Uid,
		Time:      d.Time,
		Pid:       d.Pid,
		Status:    d.Status,
		Language:  d.Language,
		TimeUsed:  d.TimeUsed,
		SpaceUsed: d.SpaceUsed,
		Code:      d.Code,
	}
}
