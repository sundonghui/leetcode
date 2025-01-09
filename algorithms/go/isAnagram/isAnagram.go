package isanagram

func isAnagram(s string, t string) bool {
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		if _, ok := sMap[t[i]]; !ok {
			return false
		}
		sMap[t[i]]--
		if sMap[t[i]] < 0 {
			return false
		}
	}
	return true
}
