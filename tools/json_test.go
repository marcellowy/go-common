// Package tools
package tools

import (
	"bytes"
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

func TestJSONUnmarshalByte(t *testing.T) {
	if !bytes.Equal(JSONMarshalByte(nil), []byte("null")) {
		t.Error("null")
		return
	}
}
