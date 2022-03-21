package logic

import (
	"context"

	"queoj/service/ask/ask"
	"queoj/service/ask/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveAskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveAskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAskLogic {
	return &RemoveAskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveAskLogic) RemoveAsk(in *ask.AskByIdReq) (*ask.Empty, error) {
	l.svcCtx.Db.Exec("delete from asks where id = ?",in.Id)
	l.svcCtx.Db.Exec("delete from replies where ask_id = ?",in.Id)
	return &ask.Empty{}, nil
}
