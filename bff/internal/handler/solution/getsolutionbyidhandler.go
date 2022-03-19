package solution

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"github.com/tal-tech/go-zero/rest/httpx"
	"queoj/bff/internal/logic/solution"
	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"
)

func GetSolutionByIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SolutionByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			advice.HandleError(w, err)
			return
		}

		l := solution.NewGetSolutionByIdLogic(r.Context(), ctx)
		resp, err := l.GetSolutionById(req)
		advice.HandleResult(w, resp, err)
	}
}
