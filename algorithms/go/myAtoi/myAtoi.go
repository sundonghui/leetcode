package myatoi

func myAtoi(s string) int {
	result, sign, i, n := 0, 1, 0, len(s)
	min, max := -1<<31, 1<<31-1
	for i < n && s[i] == ' ' {
		i++
	}
	if i >= n {
		return result
	}

	switch s[i] {
	case '-':
		sign = -1
		i++
	case '+':
		i++
	}

	for ; i < n; i++ {
		if s[i] < 48 || s[i] > 57 {
			break
		}

		result = result*10 + int(s[i]-'0')
		if result*sign > max {
			return max
		}
		if result*sign < min {
			return min
		}
	}
	return sign * result
}
