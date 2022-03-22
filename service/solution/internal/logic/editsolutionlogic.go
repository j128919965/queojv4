package logic

import (
	"context"
	"queoj/service/solution/internal/model"

	"queoj/service/solution/internal/svc"
	"queoj/service/solution/solution"

	"github.com/tal-tech/go-zero/core/logx"
)

type EditSolutionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditSolutionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditSolutionLogic {
	return &EditSolutionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditSolutionLogic) EditSolution(in *solution.SolutionDetail) (*solution.Empty, error) {
	s := &model.Solution{
		Id:       in.Id,
		Uid:      in.Uid,
		Pid:      in.Pid,
		Time:     in.Time,
		Nickname: in.Nickname,
		Title:    in.Title,
		Summary:  in.Summary,
		Content:  in.Content,
		IsTeacher: in.IsTeacher,
	}
	err := l.svcCtx.Db.Save(s).Error
	return &solution.Empty{}, err
}
