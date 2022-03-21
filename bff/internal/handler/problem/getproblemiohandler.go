package problem

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"github.com/tal-tech/go-zero/rest/httpx"
	"queoj/bff/internal/logic/problem"
	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"
)

func GetProblemIOHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Integer
		if err := httpx.Parse(r, &req); err != nil {
			advice.HandleError(w, err)
			return
		}

		l := problem.NewGetProblemIOLogic(r.Context(), ctx)
		resp, err := l.GetProblemIO(req)
		advice.HandleResult(w, resp, err)
	}
}
