package reverseleftwords

import "strings"

func reverseLeftWords(s string, n int) string {
	var left strings.Builder
	var right strings.Builder
	for i, v := range s {
		if i < n {
			left.Write([]byte(string(v)))
		} else {
			right.Write([]byte(string(v)))
		}
	}
	return right.String() + left.String()
}
