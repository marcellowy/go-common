// Package ginctx
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package ginctx

import (
	"context"
	"github.com/gin-gonic/gin"
)

func FromContext(ctx context.Context) *gin.Context {
	if newCtx, ok := ctx.(*gin.Context); ok {
		return newCtx
	}
	return nil
}
