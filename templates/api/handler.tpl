package {{.PkgName}}

import (
	"net/http"

    "github.com/j128919965/gopkg/web/advice"
	{{if .After1_1_10}}"github.com/tal-tech/go-zero/rest/httpx"{{end}}
	{{.ImportPackages}}
)

func {{.HandlerName}}(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			advice.HandleError(w,err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), ctx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}req{{end}})
		{{if .HasResp}}advice.HandleResult(w,resp,err){{else}}advice.HandleError(w,err){{end}}
	}
}
