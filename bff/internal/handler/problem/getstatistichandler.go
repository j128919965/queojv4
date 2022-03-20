package problem

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"github.com/tal-tech/go-zero/rest/httpx"
	"queoj/bff/internal/logic/problem"
	"queoj/bff/internal/svc"
)

func GetStatisticHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := problem.NewGetStatisticLogic(r.Context(), ctx)
		resp, err := l.GetStatistic()
		advice.HandleResult(w, resp, err)
	}
}
