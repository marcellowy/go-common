package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/marcellowy/go-common/gogf/vlog"
	"time"
)

// PrintResponse print response info
func PrintResponse(r *ghttp.Request) {

	n := time.Now()
	r.Middleware.Next()

	// print response status, body and cost time
	if r.Response.BufferLength() < 65536 { // 64kb
		vlog.Infof(r.GetCtx(), "response status: %d\nresponse body length: %d\nresponse body: %s\ncost: %dms",
			r.Response.Status,
			r.Response.BufferLength(),
			r.Response.BufferString(),
			time.Now().Sub(n).Milliseconds(),
		)
		return
	}

	vlog.Infof(r.GetCtx(), "response status: %d\nresponse body length: %d\ncost: %dms\n",
		r.Response.Status,
		r.Response.BufferLength(),
		time.Now().Sub(n).Milliseconds(),
	)
}
