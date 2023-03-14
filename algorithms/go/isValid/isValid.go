package isvalid

var m = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

var p = map[rune]bool{
	'(': true,
	'{': true,
	'[': true,
}

func isValid(s string) bool {
	stack := []rune{}
	overflow := []rune{}
	for _, v := range s {
		if p[v] {
			stack = append(stack, v)
			continue
		}
		if len(stack) > 0 && stack[len(stack)-1] == m[v] {
			stack = stack[0 : len(stack)-1]
			continue
		}
		overflow = append(overflow, v)
	}
	return len(overflow) == 0 && len(stack) == 0
}
