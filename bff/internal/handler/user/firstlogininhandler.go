package user

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"queoj/bff/internal/logic/user"
	"queoj/bff/internal/svc"
)

func FirstLoginInHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewFirstLoginInLogic(r.Context(), ctx)
		err := l.FirstLoginIn()
		advice.HandleError(w, err)
	}
}
