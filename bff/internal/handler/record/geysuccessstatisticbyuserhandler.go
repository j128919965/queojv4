package record

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"queoj/bff/internal/logic/record"
	"queoj/bff/internal/svc"
)

func GeySuccessStatisticByUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := record.NewGeySuccessStatisticByUserLogic(r.Context(), ctx)
		resp, err := l.GeySuccessStatisticByUser()
		advice.HandleResult(w, resp, err)
	}
}
