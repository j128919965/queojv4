// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	ask "queoj/bff/internal/handler/ask"
	message "queoj/bff/internal/handler/message"
	problem "queoj/bff/internal/handler/problem"
	record "queoj/bff/internal/handler/record"
	solution "queoj/bff/internal/handler/solution"
	user "queoj/bff/internal/handler/user"
	"queoj/bff/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login/password",
				Handler: user.LoginByPasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login/code",
				Handler: user.LoginByCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/verify/email",
				Handler: user.VerifyEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/refresh",
				Handler: user.RefreshHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/info",
				Handler: user.GetMineInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/info/email",
				Handler: user.GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/info",
				Handler: user.ChangeUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/login/first",
				Handler: user.FirstLoginInHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/password",
				Handler: user.ChangePasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/rank",
				Handler: user.GetUserRankHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/pe",
				Handler: user.PEReqHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/pe",
				Handler: user.PEGetHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/pe/all",
				Handler: user.AllPEHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/pe/app",
				Handler: user.AppPeHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/msg",
				Handler: message.GetMessageByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/msg/page",
				Handler: message.GetPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/msg/read",
				Handler: message.ReadMsgHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/msg/read/all",
				Handler: message.ReadAllHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/msg/del",
				Handler: message.DeleteByIdHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/pro",
				Handler: problem.GetProblemByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/pro/all",
				Handler: problem.GetPageHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/pro/hot",
				Handler: problem.GetHotHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/pro/tags",
				Handler: problem.GetByTagsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/pro/statistic",
				Handler: problem.GetStatisticHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/pro",
				Handler: problem.AddOrUpdateProblemHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/pro/io",
				Handler: problem.GetProblemIOHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/solution",
				Handler: solution.GetSolutionByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/solution/all",
				Handler: solution.GetAllSolutionByPidHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/solution/all/all",
				Handler: solution.AllSolutionHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPut,
				Path:    "/solution",
				Handler: solution.AddSolutionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/solution/rem",
				Handler: solution.DeleteSolutionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/solution/adm/edit",
				Handler: solution.AdminEditSolutionHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/record",
				Handler: record.SubmitRecordHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/record",
				Handler: record.GetByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/record/user",
				Handler: record.GeyByUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/record/user/statistic",
				Handler: record.GeySuccessStatisticByUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/record/status",
				Handler: record.GetStatusHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ask/all",
				Handler: ask.GetAllAskHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ask",
				Handler: ask.GetAskByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ask/replies",
				Handler: ask.GetAskRepliesHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/ask",
				Handler: ask.AddAskHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/ask/replies",
				Handler: ask.AddReplyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/ask/adm/edit",
				Handler: ask.EditAskHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/ask/adm/replies/edit",
				Handler: ask.EditReplyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/ask/amd/rem",
				Handler: ask.RemoveAskHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/ask/adm/replies/rem",
				Handler: ask.RemoveReplyHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
