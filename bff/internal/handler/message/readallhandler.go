package message

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"queoj/bff/internal/logic/message"
	"queoj/bff/internal/svc"
)

func ReadAllHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := message.NewReadAllLogic(r.Context(), ctx)
		err := l.ReadAll()
		advice.HandleError(w, err)
	}
}
