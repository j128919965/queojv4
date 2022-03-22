package user

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"queoj/bff/internal/logic/user"
	"queoj/bff/internal/svc"
)

func AllPEHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewAllPELogic(r.Context(), ctx)
		resp, err := l.AllPE()
		advice.HandleResult(w, resp, err)
	}
}
