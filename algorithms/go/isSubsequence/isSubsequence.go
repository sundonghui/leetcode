package issubsequence

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for j <= len(t) {
		if i == len(s) {
			return true
		}
		if j == len(t) {
			return false
		}
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return false
}
