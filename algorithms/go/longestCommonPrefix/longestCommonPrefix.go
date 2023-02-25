package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	min := len(strs[0])
	for _, item := range strs {
		if len(item) < min {
			min = len(item)
		}
	}
	i := 0
	for i < min {
		if isEqual(strs, i) {
			i++
			continue
		}
		break
	}
	return strs[0][0:i]
}

func isEqual(strs []string, i int) bool {
	str := strs[0][i]
	for _, item := range strs {
		if item[i] != str {
			return false
		}
	}
	return true
}
