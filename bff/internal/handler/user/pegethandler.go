package user

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"queoj/bff/internal/logic/user"
	"queoj/bff/internal/svc"
)

func PEGetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewPEGetLogic(r.Context(), ctx)
		resp, err := l.PEGet()
		advice.HandleResult(w, resp, err)
	}
}
