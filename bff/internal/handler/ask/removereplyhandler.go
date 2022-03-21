package ask

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"github.com/tal-tech/go-zero/rest/httpx"
	"queoj/bff/internal/logic/ask"
	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"
)

func RemoveReplyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AskByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			advice.HandleError(w, err)
			return
		}

		l := ask.NewRemoveReplyLogic(r.Context(), ctx)
		err := l.RemoveReply(req)
		advice.HandleError(w, err)
	}
}