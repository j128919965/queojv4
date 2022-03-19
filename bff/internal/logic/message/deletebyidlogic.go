package message

import (
	"context"
	"queoj/service/message/messageclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteByIdLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteByIdLogic {
	return DeleteByIdLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteByIdLogic) DeleteById(req types.MessageByIdReq) error {
	_, err := l.svcCtx.MessageClient.DeleteMsg(l.ctx, &messageclient.MessageByIdReq{Id: req.Id})
	return err
}
