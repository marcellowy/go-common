package vfile

import (
	"fmt"
	"io"
)

// ReadSpecifiedSize Read the specified size
// from offset read size byte
func ReadSpecifiedSize(f io.Reader, size int64) (b []byte, err error) {
	if size <= 0 || f == nil {
		return nil, fmt.Errorf("size or Reader error")
	}
	var (
		buf            = make([]byte, 0)
		readSize int64 = 0
		n              = 1024
	)
	for readSize < size {
		if (readSize + int64(n)) > size {
			n = int(size - readSize)
		}
		tmpBuf := make([]byte, n)
		var readLen int
		if readLen, err = f.Read(tmpBuf); err != nil {
			return
		}
		buf = append(buf, tmpBuf...)
		readSize += int64(readLen)
	}
	return buf, nil
}
