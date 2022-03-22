package ask

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"queoj/service/ask/askclient"
	"time"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddReplyLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddReplyLogic {
	return AddReplyLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddReplyLogic) AddReply(req types.ReplyAddReq) error {
	payLoad := l.ctx.Value("payload").(*security.PayLoad)
	userId := payLoad.UserId
	isTeacher := false
	if payLoad.Role == 2 {
		isTeacher = true
	}
	_, err := l.svcCtx.AskClient.AddReply(l.ctx,&askclient.ReplyDetail{
		AskId:    req.AskId,
		Uid:      userId,
		Time:     time.Now().Unix(),
		Nickname: req.Nickname,
		Content:  req.Content,
		IsTeacher: isTeacher,
	})
	return err
}
