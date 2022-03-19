package user

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"github.com/tal-tech/go-zero/rest/httpx"
	"queoj/bff/internal/logic/user"
	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"
)

func LoginByPasswordHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginByPasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			advice.HandleError(w, err)
			return
		}

		l := user.NewLoginByPasswordLogic(r.Context(), ctx)
		resp, err := l.LoginByPassword(req)
		advice.HandleResult(w, resp, err)
	}
}
