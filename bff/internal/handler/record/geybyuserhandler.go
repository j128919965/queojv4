package record

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"queoj/bff/internal/logic/record"
	"queoj/bff/internal/svc"
)

func GeyByUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := record.NewGeyByUserLogic(r.Context(), ctx)
		resp, err := l.GeyByUser()
		advice.HandleResult(w, resp, err)
	}
}
