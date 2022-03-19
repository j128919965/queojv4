// Code generated by goctl. DO NOT EDIT!
// Source: problem.proto

package problemclient

import (
	"context"

	"queoj/service/problem/problem"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Empty            = problem.Empty
	Integer          = problem.Integer
	ProblemDetail    = problem.ProblemDetail
	ProblemIO        = problem.ProblemIO
	ProblemPage      = problem.ProblemPage
	ProblemSynopsis  = problem.ProblemSynopsis
	ProblemsByTagReq = problem.ProblemsByTagReq
	Result           = problem.Result

	Problem interface {
		GetProblemDetail(ctx context.Context, in *Integer) (*ProblemDetail, error)
		GetAllProblem(ctx context.Context, in *Empty) (*ProblemPage, error)
		GetProblemIO(ctx context.Context, in *Integer) (*ProblemIO, error)
		GetProblemsByTags(ctx context.Context, in *ProblemsByTagReq) (*ProblemPage, error)
	}

	defaultProblem struct {
		cli zrpc.Client
	}
)

func NewProblem(cli zrpc.Client) Problem {
	return &defaultProblem{
		cli: cli,
	}
}

func (m *defaultProblem) GetProblemDetail(ctx context.Context, in *Integer) (*ProblemDetail, error) {
	client := problem.NewProblemClient(m.cli.Conn())
	return client.GetProblemDetail(ctx, in)
}

func (m *defaultProblem) GetAllProblem(ctx context.Context, in *Empty) (*ProblemPage, error) {
	client := problem.NewProblemClient(m.cli.Conn())
	return client.GetAllProblem(ctx, in)
}

func (m *defaultProblem) GetProblemIO(ctx context.Context, in *Integer) (*ProblemIO, error) {
	client := problem.NewProblemClient(m.cli.Conn())
	return client.GetProblemIO(ctx, in)
}

func (m *defaultProblem) GetProblemsByTags(ctx context.Context, in *ProblemsByTagReq) (*ProblemPage, error) {
	client := problem.NewProblemClient(m.cli.Conn())
	return client.GetProblemsByTags(ctx, in)
}
