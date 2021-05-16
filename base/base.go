package base

// IsNumber 判断一个字符串是不是全都是数字
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
