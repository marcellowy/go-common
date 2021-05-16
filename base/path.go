package base

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// PathExists 判断目录或者文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetCurrentDirectory 返回工作目录
func GetCurrentDirectory() string {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		return ""
	}

	if runtime.GOOS == "windows" {
		// 只有windows需要替换
		return strings.Replace(dir, "\\", "/", -1) //将\替换成/
	}
	return dir
}
