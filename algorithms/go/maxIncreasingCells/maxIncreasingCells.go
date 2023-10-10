package maxincreasingcells

func maxIncreasingCells(mat [][]int) int {
	type pair struct {
		x int
		y int
	}
	
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}