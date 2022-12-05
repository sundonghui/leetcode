package maxvalue

func maxValue(grid [][]int) int {
	for index1, value1 := range grid {
		for index2 := range value1 {
			if index1 > 0 && index2 > 0 {
				if grid[index1][index2-1] > grid[index1-1][index2] {
					grid[index1][index2] += grid[index1][index2-1]
				} else {
					grid[index1][index2] += grid[index1-1][index2]
				}
			} else if index1 == 0 && index2 > 0 {
				grid[index1][index2] += grid[index1][index2-1]
			} else if index1 > 0 && index2 == 0 {
				grid[index1][index2] += grid[index1-1][index2]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
