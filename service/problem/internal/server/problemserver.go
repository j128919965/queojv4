// Code generated by goctl. DO NOT EDIT!
// Source: problem.proto

package server

import (
	"context"

	"queoj/service/problem/internal/logic"
	"queoj/service/problem/internal/svc"
	"queoj/service/problem/problem"
)

type ProblemServer struct {
	svcCtx *svc.ServiceContext
}

func NewProblemServer(svcCtx *svc.ServiceContext) *ProblemServer {
	return &ProblemServer{
		svcCtx: svcCtx,
	}
}

func (s *ProblemServer) GetProblemDetail(ctx context.Context, in *problem.Integer) (*problem.ProblemDetail, error) {
	l := logic.NewGetProblemDetailLogic(ctx, s.svcCtx)
	return l.GetProblemDetail(in)
}

func (s *ProblemServer) GetAllProblem(ctx context.Context, in *problem.Empty) (*problem.ProblemPage, error) {
	l := logic.NewGetAllProblemLogic(ctx, s.svcCtx)
	return l.GetAllProblem(in)
}

func (s *ProblemServer) GetProblemIO(ctx context.Context, in *problem.Integer) (*problem.ProblemIO, error) {
	l := logic.NewGetProblemIOLogic(ctx, s.svcCtx)
	return l.GetProblemIO(in)
}

func (s *ProblemServer) GetProblemsByTags(ctx context.Context, in *problem.ProblemsByTagReq) (*problem.ProblemPage, error) {
	l := logic.NewGetProblemsByTagsLogic(ctx, s.svcCtx)
	return l.GetProblemsByTags(in)
}

func (s *ProblemServer) GetAllProblemStatistic(ctx context.Context, in *problem.Empty) (*problem.AllProblemStatistic, error) {
	l := logic.NewGetAllProblemStatisticLogic(ctx, s.svcCtx)
	return l.GetAllProblemStatistic(in)
}
