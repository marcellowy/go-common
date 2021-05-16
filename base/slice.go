package base

// RemoveSameFromStringSlice 移除 []string 中的相同元素
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
