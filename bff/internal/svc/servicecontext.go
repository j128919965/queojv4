package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/zrpc"
	"queoj/bff/internal/config"
	"queoj/service/ask/askclient"
	"queoj/service/message/messageclient"
	"queoj/service/problem/problemclient"
	"queoj/service/record/recordclient"
	"queoj/service/solution/solutionclient"
	"queoj/service/user/userclient"
)

type ServiceContext struct {
	Config         config.Config
	Redis          *redis.Redis
	UserClient     userclient.User
	ProblemClient  problemclient.Problem
	SolutionClient solutionclient.Solution
	MessageClient  messageclient.Message
	RecordClient   recordclient.Record
	AskClient      askclient.Ask
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Redis:          redis.New(c.Redis.Host,redis.WithPass(c.Redis.Pass)),
		UserClient:     userclient.NewUser(zrpc.MustNewClient(c.UserClient)),
		ProblemClient:  problemclient.NewProblem(zrpc.MustNewClient(c.ProblemClient)),
		SolutionClient: solutionclient.NewSolution(zrpc.MustNewClient(c.SolutionClient)),
		MessageClient:  messageclient.NewMessage(zrpc.MustNewClient(c.MessageClient)),
		RecordClient:   recordclient.NewRecord(zrpc.MustNewClient(c.RecordClient)),
		AskClient:      askclient.NewAsk(zrpc.MustNewClient(c.AskClient)),
	}
}
