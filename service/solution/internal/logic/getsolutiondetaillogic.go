package logic

import (
	"context"
	"queoj/service/solution/internal/model"

	"queoj/service/solution/internal/svc"
	"queoj/service/solution/solution"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetSolutionDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSolutionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSolutionDetailLogic {
	return &GetSolutionDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSolutionDetailLogic) GetSolutionDetail(in *solution.SolutionByIdReq) (*solution.SolutionDetail, error) {
	s, err := l.svcCtx.GetSolutionById(in.GetId())
	if err != nil {
		return nil, err
	}
	return ModelToDetail(s), nil
}

func ModelToDetail(s *model.Solution) *solution.SolutionDetail {
	return &solution.SolutionDetail{
		Id:       s.Id,
		Uid:      s.Uid,
		Pid:      s.Pid,
		Time:     s.Time,
		Nickname: s.Nickname,
		Title:    s.Title,
		Summary:  s.Summary,
		Content:  s.Content,
		IsTeacher: s.IsTeacher,
	}
}
