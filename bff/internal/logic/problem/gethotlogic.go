package problem

import (
	"context"
	"queoj/service/problem/problem"
	"strconv"

	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetHotLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHotLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetHotLogic {
	return GetHotLogic{
		log:    logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHotLogic) GetHot() (resp *types.ProblemHotResp, err error) {
	page, err := l.svcCtx.ProblemClient.GetAllProblem(l.ctx,&problem.Empty{})
	if err != nil {
		return nil, err
	}

	problems := map[int32]*types.ProblemHotVO{}
	for _,v := range page.Problems {
		problems[v.Id] = &types.ProblemHotVO{
			Id:    v.Id,
			Name:  v.Name,
		}
	}

	zrange, err := l.svcCtx.Redis.Zrevrange("pro:times", 0, 9)
	if err != nil {
		return nil, err
	}

	var ret []*types.ProblemHotVO
	for _, s := range zrange {
		id, _ := strconv.ParseInt(s, 10, 32)
		ret = append(ret , problems[int32(id)])
	}

	return &types.ProblemHotResp{Problems: ret},nil
}
