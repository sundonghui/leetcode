package strstr

func strStr(haystack string, needle string) int {
	for i, v := range haystack {
		right := i + len(needle)
		if right <= len(haystack) && string(needle[0]) == string(v) && needle == string(haystack[i:right]) {
			return i
		}
	}
	return -1
}
