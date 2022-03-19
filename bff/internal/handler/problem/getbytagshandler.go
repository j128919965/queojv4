package problem

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"github.com/tal-tech/go-zero/rest/httpx"
	"queoj/bff/internal/logic/problem"
	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"
)

func GetByTagsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProblemsByTagReq
		if err := httpx.Parse(r, &req); err != nil {
			advice.HandleError(w, err)
			return
		}

		l := problem.NewGetByTagsLogic(r.Context(), ctx)
		resp, err := l.GetByTags(req)
		advice.HandleResult(w, resp, err)
	}
}
