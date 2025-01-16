package simplifypath

import (
	"strings"
)

func simplifyPath(path string) string {
	stack := []string{}
	for _, dir := range strings.Split(path, "/") {
	 if dir == ".." {
		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	 } else if dir != "." && dir != "" {
	     stack = append(stack, dir)
	 }
	}
	return "/" + strings.Join(stack, "/")
}