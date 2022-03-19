package logic

import (
	"context"
	"queoj/service/message/internal/model"
	"queoj/service/message/internal/svc"
	"queoj/service/message/message"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteAllMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAllMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAllMsgLogic {
	return &DeleteAllMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAllMsgLogic) DeleteAllMsg(in *message.MessageByUserIdReq) (*message.Empty, error) {
	err := l.svcCtx.Db.Delete(&model.Message{
		Receiver: in.UserId,
	}).Error
	return &message.Empty{}, err
}
