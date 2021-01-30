package buffer

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {

	var example = []byte(`
debug: true
name: steve
`)

	v, err := New(bytes.NewBuffer(example), "yaml")
	if err != nil {
		t.Error(err)
		return
	}

	if v.GetBool("debug") != true {
		t.Error("config err")
		return
	}
}
