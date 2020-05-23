package common

import (
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//  判断一个字符串是不是全都是数字
func IsNumber(s string) bool {

	if s == "" {
		return false
	}

	var a = []byte(s)
	for _, v := range a {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

// 随机字符串
func RandStr(length int) []byte {

	var b []byte
	if length <= 0 {
		return b
	}

	var a = []byte{
		/* A-Z */
		0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48 /* 0x49, I */, 0x4A,
		0x4B, 0x4C, 0x4D, 0x4E /* 0x4F O */, 0x50, 0x51, 0x52, 0x53, 0x54,
		0x55, 0x56, 0x57, 0x58, 0x59, 0x5A,
		/* a-z */
		0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A,
		0x6B /*0x6C l */, 0x6D, 0x6E /*0x6F o*/, 0x70, 0x71, 0x72, 0x73, 0x74,
		0x75, 0x76, 0x77, 0x78, 0x79, 0x7A,
		/* 0-9 */
		/* 0x30 0 */ /* 0x31 1 */ 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
	}

	var l = len(a)
	for i := 0; i < length; i++ {
		b = append(b, a[rand.Intn(l-1)])
	}

	return b
}

// 移除 []string 中的相同元素
func RemoveSameFromStringSlice(slice []string) []string {

	var result []string
	tmp := map[string]byte{} // 存放不重复主键
	for _, s := range slice {
		l := len(tmp)
		tmp[s] = 0
		if len(tmp) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, s)
		}
	}
	return result
}

// 判断目录或者文件是否存在
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

// 将时间缀格式为YYYY-MM-DD HH:mm:ss
func FormatTime(i int) string {
	return time.Unix(int64(i), 0).Format("2006-01-02 :15:04:05")
}

// 返回工作目录
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
