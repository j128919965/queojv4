package ask

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"queoj/bff/internal/logic/ask"
	"queoj/bff/internal/svc"
)

func GetAllAskHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := ask.NewGetAllAskLogic(r.Context(), ctx)
		resp, err := l.GetAllAsk()
		advice.HandleResult(w, resp, err)
	}
}
