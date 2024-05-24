package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/marcellowy/go-common/gogf/vlog"
	"net/http"
)

// PrintRequest print request info
func PrintRequest(r *ghttp.Request) {

	// if request method is get print method and request rui
	// if request method is post print content include body
	switch r.Method {
	case http.MethodPost, http.MethodGet:
		vlog.Infof(r.GetCtx(), "%s %s", r.Method, r.RequestURI)
	}

	requestBody := r.GetBody()
	if len(requestBody) > 0 {
		var requestBodyContent = []byte("request body too large, more than 65535 byte")
		if len(requestBody) < 65536 {
			requestBodyContent = requestBody
		}
		vlog.Infof(r.GetCtx(), "request body length: %d\nreqeust body: %s", len(requestBody), requestBodyContent)
	}

	r.Middleware.Next()
}
