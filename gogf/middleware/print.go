// Package middleware
package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/marcellowy/go-common/gogf/vlog"
	"net/http"
	"time"
)

// Print print request info
func Print(r *ghttp.Request) {

	n := time.Now()

	// if request method is get print method and request rui
	// if request method is post print content include body
	switch r.Method {
	case http.MethodPost, http.MethodGet:
		vlog.Infof(r.GetCtx(), "%s %s", r.Method, r.RequestURI)
	}

	requestBody := r.GetBody()
	if len(requestBody) > 0 {
		vlog.Infof(r.GetCtx(), "request body length: %d\nreqeust body: %s", len(requestBody), requestBody)
	}

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
