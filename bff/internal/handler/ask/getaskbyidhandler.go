package ask

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"github.com/tal-tech/go-zero/rest/httpx"
	"queoj/bff/internal/logic/ask"
	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"
)

func GetAskByIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AskByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			advice.HandleError(w, err)
			return
		}

		l := ask.NewGetAskByIdLogic(r.Context(), ctx)
		resp, err := l.GetAskById(req)
		advice.HandleResult(w, resp, err)
	}
}
