package logic

import (
	"context"
	"queoj/service/solution/internal/model"
	"time"

	"queoj/service/solution/internal/svc"
	"queoj/service/solution/solution"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddSolutionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSolutionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSolutionLogic {
	return &AddSolutionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddSolutionLogic) AddSolution(in *solution.SolutionAddReq) (*solution.Empty, error) {
	smRune := []rune(in.Content)
	sm := len(smRune)
	if sm > 80 {
		sm = 80
	}
	summary := string(smRune[:sm])
	s := &model.Solution{
		Uid:      in.Uid,
		Pid:      in.Pid,
		Time:     uint64(time.Now().Unix()),
		Nickname: in.Nickname,
		Title:    in.Title,
		Summary:  summary,
		Content:  in.Content,
	}

	err := l.svcCtx.Db.Create(s).Error
	if err != nil {
		return nil, err
	}

	return &solution.Empty{}, nil
}
