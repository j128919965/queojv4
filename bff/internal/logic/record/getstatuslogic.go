package record

import (
	"context"
	"queoj/service/record/recordclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetStatusLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetStatusLogic {
	return GetStatusLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStatusLogic) GetStatus(req types.RecordByIdReq) (resp *types.RecordStatus, err error) {
	id, err := l.svcCtx.RecordClient.GetRecordStatus(l.ctx, &recordclient.RecordByIdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.RecordStatus{Status: id.Status},nil
}
