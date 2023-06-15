// Package middleware
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/marcellowy/go-common/consts"
	"github.com/marcellowy/go-common/server/ginctx"
)

func TraceId(ctx *gin.Context) {
	if val, ok := ctx.Get(consts.DefaultTraceKey); val == nil || !ok {
		ctx = ginctx.AddTrace(ctx).(*gin.Context)
	}
	ctx.Next()
}
