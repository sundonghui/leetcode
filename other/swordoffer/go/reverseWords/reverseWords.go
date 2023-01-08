package reversewords

import (
	"bytes"
	"strings"
)

func reverseWords(s string) string {
	list := strings.Split(s, " ")
	tmpList := make([]string, 0, len(list))
	for _, item := range list {
		if item != "" {
			tmpList = append([]string{item}, tmpList...)
		}
	}

	var sb bytes.Buffer
	for index, item := range tmpList {
		sb.Write([]byte(item))
		if index < len(tmpList)-1 {
			sb.Write([]byte(` `))
		}
	}

	return strings.Trim(sb.String(), " ")
}
