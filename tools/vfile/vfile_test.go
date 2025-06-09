package vfile

import (
	"os"
	"testing"
)

func TestReadSpecifiedSize(t *testing.T) {

	c := make([]byte, 0)
	for i := 0; i < 10240; i++ {
		c = append(c, []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")...)
	}

	filename := "aa.txt"
	readSize := len(c)
	err := os.WriteFile(filename, c, os.ModePerm)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if err = os.RemoveAll(filename); err != nil {

		}
	}()

	var ff *os.File
	if ff, err = os.Open(filename); err != nil {
		t.Error(err)
		return
	}
	defer func() {
		_ = ff.Close()
	}()

	var buf []byte
	if buf, err = ReadSpecifiedSize(ff, int64(readSize)); err != nil {
		t.Error(err)
		return
	}

	if len(buf) != readSize {
		t.Error("len err")
	}
}
