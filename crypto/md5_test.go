package crypto

import (
	"testing"
)

func TestFileMd5(t *testing.T) {
	if _, err := FileMd5("H:/tmp/config.yaml"); err != nil {
		t.Error(err)
	}
}
