package file

import "testing"

func TestNew(t *testing.T) {
	c, err := New("H:/tmp/config.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	b := c.GetBool("debug")
	if b == false {
		t.Errorf("bool value error")
		return
	}

}
