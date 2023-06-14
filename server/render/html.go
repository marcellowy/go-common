// Package render
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package render

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/marcellowy/go-common/config"
	"github.com/marcellowy/go-common/server/ginctx"
	"github.com/marcellowy/go-common/tools"
	"net/http"
)

type renderData struct {
	Key   string
	Value string
}

// HTML render
func HTML(ctx context.Context, name string, obj any) {
	newCtx := ginctx.FromContext(ctx)
	var viewerData []*renderData
	if err := json.Unmarshal(tools.JSONMarshalByte(config.Get("viewer.data")), &viewerData); err != nil {
		return
	}

	if data, ok := obj.(gin.H); ok {
		// add some view
		for _, v := range viewerData {
			data[v.Key] = v.Value
		}
		newCtx.HTML(http.StatusOK, name, data)
		return
	}

	if data, ok := obj.(map[string]any); ok {
		// add some view
		for _, v := range viewerData {
			data[v.Key] = v.Value
		}
		newCtx.HTML(http.StatusOK, name, data)
		return
	}
	newCtx.HTML(http.StatusOK, name, obj)
}
