package logic

import (
	"context"
	"fmt"
	"queoj/service/message/internal/model"

	"queoj/service/message/internal/svc"
	"queoj/service/message/message"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMsgLogic {
	return &DeleteMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMsgLogic) DeleteMsg(in *message.MessageByIdReq) (*message.Empty, error) {
	key := fmt.Sprintf("msg:id:%d", in.Id)
	l.svcCtx.Redis.Del(key)
	err := l.svcCtx.Db.Delete(&model.Message{Id: in.Id}).Error
	return &message.Empty{}, err
}
