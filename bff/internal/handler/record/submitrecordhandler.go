package record

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	"github.com/tal-tech/go-zero/rest/httpx"
	"queoj/bff/internal/logic/record"
	"queoj/bff/internal/svc"
	"queoj/bff/internal/types"
)

func SubmitRecordHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecordAddReq
		if err := httpx.Parse(r, &req); err != nil {
			advice.HandleError(w, err)
			return
		}

		l := record.NewSubmitRecordLogic(r.Context(), ctx)
		resp, err := l.SubmitRecord(req)
		advice.HandleResult(w, resp, err)
	}
}
