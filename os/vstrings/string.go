package vstrings

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"strings"
)

// ReplaceLineBreaks 移除\r, \r\n, \n,
func ReplaceLineBreaks(s string) string {
	r := strings.NewReplacer("\r", "", "\n", "")
	return r.Replace(s)
}

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
