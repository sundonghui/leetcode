package romantoint

func romanToInt(s string) int {
	var strToInt = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	if len(s) <= 1 {
		return strToInt[s]
	}
	res := 0
	for i := 0; i < len(s)-1; i++ {
		cur := strToInt[string(s[i])]
		next := strToInt[string(s[i+1])]
		if cur >= next {
			res += cur
		} else {
			res -= cur
		}
	}
	res += strToInt[string(s[len(s)-1])]
	return res
}
