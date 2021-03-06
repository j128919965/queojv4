// Code generated by goctl. DO NOT EDIT!
// Source: ask.proto

package server

import (
	"context"

	"queoj/service/ask/ask"
	"queoj/service/ask/internal/logic"
	"queoj/service/ask/internal/svc"
)

type AskServer struct {
	svcCtx *svc.ServiceContext
}

func NewAskServer(svcCtx *svc.ServiceContext) *AskServer {
	return &AskServer{
		svcCtx: svcCtx,
	}
}

func (s *AskServer) GetAskById(ctx context.Context, in *ask.AskByIdReq) (*ask.AskDetail, error) {
	l := logic.NewGetAskByIdLogic(ctx, s.svcCtx)
	return l.GetAskById(in)
}

func (s *AskServer) GetAllAsk(ctx context.Context, in *ask.Empty) (*ask.AskList, error) {
	l := logic.NewGetAllAskLogic(ctx, s.svcCtx)
	return l.GetAllAsk(in)
}

func (s *AskServer) GetReplyByAskId(ctx context.Context, in *ask.AskByIdReq) (*ask.ReplyList, error) {
	l := logic.NewGetReplyByAskIdLogic(ctx, s.svcCtx)
	return l.GetReplyByAskId(in)
}

func (s *AskServer) AddAsk(ctx context.Context, in *ask.AskDetail) (*ask.Empty, error) {
	l := logic.NewAddAskLogic(ctx, s.svcCtx)
	return l.AddAsk(in)
}

func (s *AskServer) EditAsk(ctx context.Context, in *ask.AskDetail) (*ask.Empty, error) {
	l := logic.NewEditAskLogic(ctx, s.svcCtx)
	return l.EditAsk(in)
}

func (s *AskServer) AddReply(ctx context.Context, in *ask.ReplyDetail) (*ask.Empty, error) {
	l := logic.NewAddReplyLogic(ctx, s.svcCtx)
	return l.AddReply(in)
}

func (s *AskServer) EditReply(ctx context.Context, in *ask.ReplyDetail) (*ask.Empty, error) {
	l := logic.NewEditReplyLogic(ctx, s.svcCtx)
	return l.EditReply(in)
}

func (s *AskServer) RemoveAsk(ctx context.Context, in *ask.AskByIdReq) (*ask.Empty, error) {
	l := logic.NewRemoveAskLogic(ctx, s.svcCtx)
	return l.RemoveAsk(in)
}

func (s *AskServer) RemoveReply(ctx context.Context, in *ask.ReplyByIdReq) (*ask.Empty, error) {
	l := logic.NewRemoveReplyLogic(ctx, s.svcCtx)
	return l.RemoveReply(in)
}
