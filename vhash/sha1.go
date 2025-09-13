package vhash

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

func Sha1(s string) string {
	sum := sha1.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func Sha1File(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()
	h := sha1.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
