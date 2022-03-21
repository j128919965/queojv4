package solution

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"github.com/tal-tech/go-zero/rest/httpx"
	"queoj/bff/internal/logic/solution"
	"queoj/bff/internal/svc"
)

func AllSolutionHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := solution.NewAllSolutionLogic(r.Context(), ctx)
		resp, err := l.AllSolution()
		advice.HandleResult(w, resp, err)
	}
}
