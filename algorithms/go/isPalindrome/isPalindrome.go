package ispalindrome

import "strings"

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

func IsPalindrome(s string) bool {
	var sgood string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			sgood += string(s[i])
		}
	}

	n := len(sgood)
	sgood = strings.ToLower(sgood)
	for i := 0; i < n/2; i++ {
		if sgood[i] != sgood[n-1-i] {
			return false
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') ||
		(ch >= 'a' && ch <= 'z') ||
		(ch >= '0' && ch <= '9')
}
