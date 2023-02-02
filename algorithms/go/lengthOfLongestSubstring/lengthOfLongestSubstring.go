package lengthoflongestsubstring

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int, len(s))
	l, r, max := 0, 0, 0
	for r < len(s) {
		vr := s[r]
		if m[vr] > 0 {
			lv := s[l]
			l++
			m[lv]--
			continue
		}
		if r-l+1 > max {
			max = r - l + 1
		}
		m[vr]++
		r++
	}
	return max
}
