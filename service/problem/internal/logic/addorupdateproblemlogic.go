package logic

import (
	"context"
	"queoj/service/problem/internal/model"

	"queoj/service/problem/internal/svc"
	"queoj/service/problem/problem"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddOrUpdateProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrUpdateProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrUpdateProblemLogic {
	return &AddOrUpdateProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddOrUpdateProblemLogic) AddOrUpdateProblem(in *problem.AddOrUpdateProblemReq) (*problem.Empty, error) {
	if in.Detail.Id != 0{
		pro, tags := DetailToProblem(in.Detail)
		l.svcCtx.Db.Exec("delete from problem_tags where pid = ?",pro.Id)
		l.svcCtx.Db.Save(pro)
		for _, tag := range tags {
			l.svcCtx.Db.Create(tag)
		}
	}else {
		pro, tags := DetailToProblem(in.Detail)
		l.svcCtx.Db.Create(pro)
		for _, tag := range tags {
			tag.Pid = pro.Id
			l.svcCtx.Db.Create(tag)
		}
		in.Io.Pid = pro.Id
	}

	// save IO
	l.svcCtx.Db.Exec("delete from problem_ios where pid = ?",in.Io.Pid)
	l.svcCtx.Db.Create(ProtoIOToModelIO(in.Io))
	return &problem.Empty{}, nil
}

func ProtoIOToModelIO (io *problem.ProblemIO) *model.ProblemIO {
	return &model.ProblemIO{
		Pid:    io.Pid,
		InTxt:  io.InTxt,
		OutTxt: io.OutTxt,
	}
}