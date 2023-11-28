package reversewords

func reverseWords(s string) (res string) {
	s = " " + s + " "
	l, r := len(s)-1, len(s)-1
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == ' ' {
			l, r = i, l
			if r > l+1 {
				res = res + s[l+1:r] + " "
			}
		}
	}
	return res[:len(res)-1]
}

func ReverseWords(s string) string {
	ret := []byte{}
	for i := 0; i < len(s); {
		start := i
		for i < len(s) && s[i] != ' ' {
			i++
		}
		for p := start; p < i; p++ {
			pointer := start + i - 1 - p
			ret = append(ret, s[pointer])
		}
		for i < len(s) && s[i] == ' ' {
			i++
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}
