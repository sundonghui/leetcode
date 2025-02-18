package solve

// 从边缘开始传染，未被传染的就是被包围的
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		dfs(board, i, 0)
		dfs(board, i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs(board, 0, i)
		dfs(board, m-1, i)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) ||
		j < 0 || j >= len(board[0]) ||
		board[i][j] != 'O' {
		return
	}
	board[i][j] = 'A'
	dfs(board, i-1, j)
	dfs(board, i+1, j)
	dfs(board, i, j-1)
	dfs(board, i, j+1)
}
