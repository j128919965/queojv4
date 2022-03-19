// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

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
		},
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
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPut,
				Path:    "/solution",
				Handler: solution.AddSolutionHandler(serverCtx),
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
				Path:    "/record/status",
				Handler: record.GetStatusHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
