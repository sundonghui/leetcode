package convert

import "bytes"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	matrix := make([][]byte, numRows)
	t, x := numRows*2-2, 0
	for i, ch := range s {
		matrix[x] = append(matrix[x], byte(ch))
		if i%t < numRows-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(matrix, nil))
}
