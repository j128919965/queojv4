package record

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"queoj/service/record/recordclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GeyByUserLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGeyByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GeyByUserLogic {
	return GeyByUserLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GeyByUserLogic) GeyByUser() (resp *types.RecordList, err error) {
	user := l.ctx.Value("payload").(*security.PayLoad).UserId
	records, err := l.svcCtx.RecordClient.GetRecordByUser(l.ctx, &recordclient.RecordByUserReq{UserId:user})
	if err != nil {
		return nil, err
	}

	var ret []*types.RecordDetail
	for _, record := range records.Records {
		ret = append(ret,ProroToDetail(record))
	}

	return &types.RecordList{Records: ret},nil
}
