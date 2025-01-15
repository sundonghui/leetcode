package uniquepaths

func uniquePaths(m int, n int) int {
    depts := make([][]int, m)
	for i := range depts {
		tempArr := make([]int, n)
		tempArr[0] = 1
		depts[i] = tempArr
	}
	for i := 0; i < n; i++ {
		depts[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			depts[i][j] = depts[i-1][j] + depts[i][j-1]
		}
	}
	return depts[m-1][n-1]
}