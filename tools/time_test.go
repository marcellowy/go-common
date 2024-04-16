package tools

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	a := FormatTime(time.Now().Unix(), 0)
	// output: xxxx-xx-xx xx:xx:xx
	fmt.Println(a)
}
