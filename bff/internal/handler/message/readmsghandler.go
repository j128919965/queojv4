package message

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"github.com/tal-tech/go-zero/rest/httpx"
	"queoj/bff/internal/logic/message"
	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"
)

func ReadMsgHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			advice.HandleError(w, err)
			return
		}

		l := message.NewReadMsgLogic(r.Context(), ctx)
		err := l.ReadMsg(req)
		advice.HandleError(w, err)
	}
}
