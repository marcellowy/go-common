package vlog

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestNew(t *testing.T) {
	var (
		glog = New("aa")
		ctx  = context.Background()
	)
	glog.Info(ctx, "1")
	glog.Info(ctx, "2")
	glog.Error(ctx, "error")
	g.Log("aa").Skip(0).Line().Info(ctx, "================")
	Info(ctx, "11111111")
}
