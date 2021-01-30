package crypto

import (
	"testing"
)

func TestFileMd5(t *testing.T) {

}

func TestMd5(t *testing.T) {
	if Md5([]byte(`a`)) != "0cc175b9c0f1b6a831c399e269772661" {
		t.Error("md5 err")
	}
}
