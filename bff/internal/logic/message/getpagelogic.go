package message

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"queoj/service/message/messageclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetPageLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetPageLogic {
	return GetPageLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPageLogic) GetPage(req types.PageReq) (resp *types.MessageaPageResp, err error) {
	r := &messageclient.PageRequest{
		PageSize:   req.PageSize,
		PageNumber: req.PageNumber,
		LastId:     req.LastId,
		HasCount:   req.HasCount,
	}
	page, err := l.svcCtx.MessageClient.GetMessagePage(l.ctx, &messageclient.MessageReqByReceiver{
		Receiver: l.ctx.Value("payload").(*security.PayLoad).UserId,
		Noread:   nil,
		Page:     r,
	})
	if err != nil {
		return nil, err
	}
	var ret []*types.Message
	for _, m := range page.Messages {
		ret = append(ret, &types.Message{
			Id:       m.Id,
			Receiver: m.Receiver,
			Time:     m.Time,
			Read:     m.Read,
			Type:     byte(m.Type),
			Title:    m.Title,
			Content:  m.Content,
		})
	}

	return &types.MessageaPageResp{
		Messages:   ret,
		TotalCount: &page.TotalCount,
	},nil
}
