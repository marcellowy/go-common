package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/marcellowy/go-common/gogf/vlog"
	"net/http"
)

// PrintRequest print request info
func PrintRequest(r *ghttp.Request) {
	// if request method is GET print method and request rui
	// if request method is POST print request body
	switch r.Method {
	case http.MethodPost, http.MethodGet:
		vlog.Infof(r.GetCtx(), "%s %s", r.Method, r.RequestURI)
		requestBody := r.GetBody()
		if len(requestBody) < DefaultPrintRequestBodyMaxSize {
			vlog.Infof(r.GetCtx(), "request body length: %d\nreqeust body: %s", len(requestBody), requestBody)
		} else {
			vlog.Infof(r.GetCtx(), "request body length: %d\nreqeust body: %s", len(requestBody), "request body too large")
		}
	}
	r.Middleware.Next()
}
