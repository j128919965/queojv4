package record

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"queoj/service/record/recordclient"
	"time"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubmitRecordLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) SubmitRecordLogic {
	return SubmitRecordLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitRecordLogic) SubmitRecord(req types.RecordAddReq) (resp *types.Long, err error) {
	user := l.ctx.Value("payload").(*security.PayLoad).UserId
	record := &recordclient.RecordDetail{
		Uid:      user,
		Time:     uint64(time.Now().Unix()),
		Pid:      req.Pid,
		Language: req.Language,
		Code:     req.Code,
	}
	idReq, err := l.svcCtx.RecordClient.AddRecord(l.ctx, record)
	if err != nil {
		return nil, err
	}
	return &types.Long{Id: idReq.Id},nil
}
