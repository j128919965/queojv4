package logic

import (
	"context"
	"queoj/service/ask/internal/model"

	"queoj/service/ask/ask"
	"queoj/service/ask/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetReplyByAskIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetReplyByAskIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReplyByAskIdLogic {
	return &GetReplyByAskIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetReplyByAskIdLogic) GetReplyByAskId(in *ask.AskByIdReq) (*ask.ReplyList, error) {
	rs := l.svcCtx.GetRepliesByAsk(in.Id)
	var ret []*ask.ReplyDetail
	for _, r := range rs {
		ret = append(ret,ModelReplyToProtoReply(r))
	}
	return &ask.ReplyList{Replies: ret}, nil
}

func ModelReplyToProtoReply(r *model.Reply) *ask.ReplyDetail {
	return &ask.ReplyDetail{
		Id:       r.Id,
		AskId:    r.AskId,
		Uid:      r.Uid,
		Time:     r.Time,
		Nickname: r.Nickname,
		Content:  r.Content,
	}
}