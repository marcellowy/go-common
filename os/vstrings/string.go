package vstrings

import "strings"

// ReplaceLineBreaks 移除\r, \r\n, \n,
func ReplaceLineBreaks(s string) string {
	r := strings.NewReplacer("\r", "", "\n", "")
	return r.Replace(s)
}
