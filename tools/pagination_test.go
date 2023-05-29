// Package tools
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package tools

import "testing"

func TestPagination(t *testing.T) {
	p := Pagination(100, 1, 2, 5)
	if p.TotalPage != 50 {
		t.Error("total page err")
		return
	}
	if len(p.Pages) != 5 {
		t.Error("pages err")
		return
	}
	if p.NextPage == 0 {
		t.Error("next page err")
		return
	}
}
