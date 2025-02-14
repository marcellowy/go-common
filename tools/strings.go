package tools

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"strings"
)

// GBKToUTF8 GBKToUTF8
func GBKToUTF8(s string) string {
	var err error
	reader := transform.NewReader(strings.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	var sByte []byte
	if sByte, err = io.ReadAll(reader); err == nil {
		s = string(sByte)
	}
	return s
}

func UTF8ToGBK(s string) string {
	reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewEncoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return s
	}
	return string(d)
}
