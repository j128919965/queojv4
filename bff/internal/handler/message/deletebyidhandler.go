package message

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"github.com/tal-tech/go-zero/rest/httpx"
	"queoj/bff/internal/logic/message"
	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"
)

func DeleteByIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			advice.HandleError(w, err)
			return
		}

		l := message.NewDeleteByIdLogic(r.Context(), ctx)
		err := l.DeleteById(req)
		advice.HandleError(w, err)
	}
}
