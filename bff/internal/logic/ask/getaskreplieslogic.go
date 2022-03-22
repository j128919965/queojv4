package ask

import (
	"context"
	"queoj/service/ask/askclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAskRepliesLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAskRepliesLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAskRepliesLogic {
	return GetAskRepliesLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAskRepliesLogic) GetAskReplies(req types.AskByIdReq) (resp *types.ReplyList, err error) {
	rs, err := l.svcCtx.AskClient.GetReplyByAskId(l.ctx, &askclient.AskByIdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	var ret []*types.ReplyDetail
	for _, r := range rs.Replies {
		ret = append(ret , &types.ReplyDetail{
			Id:       r.Id,
			AskId:    r.AskId,
			Uid:      r.Uid,
			Time:     r.Time,
			Nickname: r.Nickname,
			Content:  r.Content,
			IsTeacher: r.IsTeacher,
		})
	}
	return &types.ReplyList{Replies: ret},nil
}
