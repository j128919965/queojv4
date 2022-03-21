package ask

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"github.com/tal-tech/go-zero/rest/httpx"
	"queoj/bff/internal/logic/ask"
	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"
)

func EditReplyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReplyDetail
		if err := httpx.Parse(r, &req); err != nil {
			advice.HandleError(w, err)
			return
		}

		l := ask.NewEditReplyLogic(r.Context(), ctx)
		err := l.EditReply(req)
		advice.HandleError(w, err)
	}
}
