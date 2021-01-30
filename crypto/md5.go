package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// Md5 计算md5值
func Md5(s []byte) string {
	m := md5.New()
	m.Write(s)
	return hex.EncodeToString(m.Sum(nil))
}

// FileMd5 计算文件md5值
func FileMd5(file string) (string, error) {

	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = f.Close()
	}()

	m := md5.New()
	if _, err = io.Copy(m, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(m.Sum(nil)), nil
}
