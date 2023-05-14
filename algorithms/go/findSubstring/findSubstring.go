package findsubstring

func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	matchLen := wordLen * len(words)
	if len(s) < matchLen {
		return []int{}
	}
	m := make(map[string]int, len(words))
	for _, v := range words {
		m[v]++
	}

	res := []int{}
	i := 0
	for i <= len(s)-wordLen*len(words) {
		if m[string(s[i:i+wordLen])] > 0 {
			if isMatch(s[i:i+matchLen], words, wordLen) {
				res = append(res, i)
			}
		}
		i++
	}
	return res
}

func isMatch(s string, words []string, l int) bool {
	m := make(map[string]int, len(words))
	for _, v := range words {
		m[v]++
	}
	for i := 0; i < len(s); {
		tmp := s[i : i+l]
		if m[tmp] > 0 {
			m[tmp]--
		} else {
			return false
		}
		i += l
	}
	return true
}
