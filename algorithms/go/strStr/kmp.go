package strstr

func kmp(haystack, needle string) int {
	if len(needle) == 0 {
		return -1
	}

	next := make([]int, len(needle))
	buildNext(next, needle)
	j := -1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j]
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == len(needle)-1 {
			return i - len(needle) + 1
		}
	}
	return -1
}

func buildNext(next []int, needle string) {
	j := -1
	next[0] = j
	for i := 1; i < len(needle); i++ {
		for j >= 0 && needle[i] != needle[j+1] {
			j = next[j]
		}
		if needle[i] == needle[j+1] {
			j++
		}
		next[i] = j
	}
}
