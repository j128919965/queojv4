package problem

import (
	"context"
	"queoj/service/problem/problemclient"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddOrUpdateProblemLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrUpdateProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddOrUpdateProblemLogic {
	return AddOrUpdateProblemLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrUpdateProblemLogic) AddOrUpdateProblem(req types.ProblemAddOrUpdateDto) error {
	in := problemclient.AddOrUpdateProblemReq{
		Detail: &problemclient.ProblemDetail{
			Id:          req.Detail.Id,
			Name:        req.Detail.Name,
			Point:       req.Detail.Point,
			Level:       req.Detail.Level,
			Description: req.Detail.Description,
			Tags:        req.Detail.Tags,
			TimeLimit:   req.Detail.TimeLimit,
			SpaceLimit:  req.Detail.SpaceLimit,
		},
		Io:     &problemclient.ProblemIO{
			Pid:    req.Detail.Id,
			InTxt:  req.InText,
			OutTxt: req.OutText,
		},
	}
	_, err := l.svcCtx.ProblemClient.AddOrUpdateProblem(l.ctx,&in)
	if err != nil {
		return err
	}
	return nil
}
