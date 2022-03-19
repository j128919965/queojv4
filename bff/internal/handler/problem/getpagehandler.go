package problem

import (
	"net/http"

	"github.com/j128919965/gopkg/web/advice"
	_ "github.com/tal-tech/go-zero/rest/httpx"
	"queoj/bff/internal/logic/problem"
	"queoj/bff/internal/svc"
)

func GetPageHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := problem.NewGetPageLogic(r.Context(), ctx)
		resp, err := l.GetPage()
		advice.HandleResult(w, resp, err)
	}
}
