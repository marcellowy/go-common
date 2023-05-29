// Package vlog
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package vlog

import (
	"context"
	"testing"
)

func TestInfo(t *testing.T) {
	Debug(context.TODO(), "")
	Debugf(context.TODO(), "")
	Info(context.TODO(), "")
	Infof(context.TODO(), "")
	Warning(context.TODO(), "")
	Warningf(context.TODO(), "")
	Error(context.TODO(), "")
	Errorf(context.TODO(), "")
}
