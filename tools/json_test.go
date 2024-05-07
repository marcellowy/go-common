// Package tools
package tools

import (
	"testing"
)

func TestJSONMarshalString(t *testing.T) {

	var a = make(map[string]string)
	a["name"] = "json"
	var s = JSONMarshalString(a)
	if s != "{\"name\":\"json\"}" {
		t.Error("marshal err")
		return
	}
}
