package firstuniqchar

func firstUniqChar(s string) byte {
	m := make(map[byte]int, len(s))
	bytes := []byte(s)
	for _, v := range bytes {
		m[v]++
	}
	for _, v := range bytes {
		if m[v] == 1 {
			return v
		}
	}
	return ' '
}
