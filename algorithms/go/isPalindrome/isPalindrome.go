package ispalindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	origin := x
	r := 0
	for x >= 10 {
		tmp := x % 10
		r = r*10 + tmp
		x /= 10
	}
	if x == 0 {
		return false
	}
	r = r*10 + x
	return r == origin
}
