package setzeroes

func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, list := range matrix {
		for j, v := range list {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, list := range matrix {
		for j := range list {
			if row[i] || col[j] {
				list[j] = 0
			}
		}
	}
}
