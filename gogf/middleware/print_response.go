package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/marcellowy/go-common/gogf/vlog"
	"net/http"
	"time"
)

// PrintResponse print response info
func PrintResponse(r *ghttp.Request) {
	n := time.Now()
	r.Middleware.Next()
	switch r.Method {
	case http.MethodPost, http.MethodGet:
		// print response status, body and cost time
		if r.Response.BufferLength() < DefaultPrintResponseBodyMaxSize { // 64kb
			vlog.Infof(r.GetCtx(), "response status: %d\nresponse body length: %d\nresponse body: %s\ncost: %dms",
				r.Response.Status,
				r.Response.BufferLength(),
				r.Response.BufferString(),
				time.Now().Sub(n).Milliseconds(),
			)
		} else {
			vlog.Infof(r.GetCtx(), "response status: %d\nresponse body length: %d\ncost: %dms\n",
				r.Response.Status,
				r.Response.BufferLength(),
				time.Now().Sub(n).Milliseconds(),
			)
		}
	}

}
