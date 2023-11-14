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

func longestCommonPrefix_(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	base := strs[0]
	for i := range base {
		for _, str := range strs[1:] {
			if i >= len(str) || str[i] != base[i] {
				return base[:i]
			}
		}
	}
	return base
}
