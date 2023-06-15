// Package ginctx
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package ginctx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/marcellowy/go-common/consts"
	"github.com/marcellowy/go-common/tools"
)

func New() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, consts.DefaultTraceKey,
		tools.RandomString(16, tools.RandomLowercase|tools.RandomDigital|tools.RandomMajuscule))
	return ctx
}

func AddTrace(ctx context.Context) context.Context {
	s := tools.RandomString(16, tools.RandomLowercase|tools.RandomDigital|tools.RandomMajuscule)
	if newCtx, ok := ctx.(*gin.Context); ok {
		newCtx.Set(consts.DefaultTraceKey, s)
	} else {
		ctx = context.WithValue(ctx, consts.DefaultTraceKey, s)
	}
	return ctx
}

func ReadTrace(ctx context.Context) string {
	if newCtx, ok := ctx.(*gin.Context); ok {
		if value, exists := newCtx.Get(consts.DefaultTraceKey); exists {
			if id, strOk := value.(string); strOk {
				return id
			}
		}
	} else {
		if id, traceOk := ctx.Value(consts.DefaultTraceKey).(string); traceOk {
			return id
		}
	}
	return ""
}

func FromContext(ctx context.Context) *gin.Context {
	if newCtx, ok := ctx.(*gin.Context); ok {
		return newCtx
	}
	return nil
}
