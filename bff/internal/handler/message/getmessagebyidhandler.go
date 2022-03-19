package message

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"github.com/tal-tech/go-zero/rest/httpx"
	"queoj/bff/internal/logic/message"
	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"
)

func GetMessageByIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			advice.HandleError(w, err)
			return
		}

		l := message.NewGetMessageByIdLogic(r.Context(), ctx)
		resp, err := l.GetMessageById(req)
		advice.HandleResult(w, resp, err)
	}
}
