// Package render
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package render

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/marcellowy/go-common/config"
	"github.com/marcellowy/go-common/server/ginctx"
	"net/http"
)

// HTML render
func HTML(ctx context.Context, name string, obj any) {
	newCtx := ginctx.FromContext(ctx)
	var viewerData = config.GetStringMap("viewer.data")
	if data, ok := obj.(gin.H); ok {
		// add some view
		for k, v := range viewerData {
			data[k] = v
		}
		newCtx.HTML(http.StatusOK, name, data)
		return
	}

	if data, ok := obj.(map[string]any); ok {
		// add some view
		for k, v := range viewerData {
			data[k] = v
		}
		newCtx.HTML(http.StatusOK, name, data)
		return
	}
	newCtx.HTML(http.StatusOK, name, obj)
}
