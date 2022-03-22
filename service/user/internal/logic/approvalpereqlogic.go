package logic

import (
	"context"
	"queoj/service/message/messageclient"
	"queoj/service/user/internal/model"
	"queoj/service/user/internal/svc"
	"queoj/service/user/user"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type ApprovalPEReqLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApprovalPEReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApprovalPEReqLogic {
	return &ApprovalPEReqLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ApprovalPEReqLogic) ApprovalPEReq(in *user.ApprovalPrivilegeEscalationReq) (*user.Result, error) {
	var p model.PeReq
	l.svcCtx.Db.Where("id = ?",in.Id).Find(&p)
	p.Approval = in.Approval
	l.svcCtx.Db.Save(p)

	if in.Approval == 1 {
		l.svcCtx.Db.Exec("update users set role = ? where id = ?",p.Role,p.UserId)
		l.svcCtx.MessageClient.SendMessage(l.ctx , &messageclient.MessageDto{
			Receiver:  p.UserId,
			Time:      time.Now().Unix(),
			Title:     "权限审批已通过",
			Content:   "您已获得权限 ： "+GetRoleName(p.Role)+"。\n重新登录后权限生效。",
		})
	}else{
		l.svcCtx.MessageClient.SendMessage(l.ctx , &messageclient.MessageDto{
			Receiver:  p.UserId,
			Time:      time.Now().Unix(),
			Title:     "权限审批已驳回",
			Content:   "管理员已驳回您申请的 ："+GetRoleName(p.Role)+"权限。",
		})
	}

	return &user.Result{}, nil
}

func GetRoleName (role uint32)string{
	switch role {
	case 1: return "学生"
	case 2: return "教师"
	default: return "管理员"
	}
}