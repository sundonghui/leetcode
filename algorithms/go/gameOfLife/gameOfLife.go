package gameoflife

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	//原地算法的话，要引入新的变量
	//2--原来是0，但是变成1
	//3--原来是1，但是变成0

	//八个方向
	dirx := [8]int{0, 0, 1, -1, 1, 1, -1, -1}
	diry := [8]int{1, -1, 0, 0, 1, -1, 1, -1}

	valid := func(i, j int) bool {
		if i >= 0 && i < m && j >= 0 && j < n {
			return true
		}
		return false
	}

	count := func(i, j int) int {
		//计算活细胞的个数
		cnt := 0
		//对八个方向找
		for k := 0; k < 8; k++ {
			newx := i + dirx[k]
			newy := j + diry[k]
			//如果合法且是
			if valid(newx, newy) && (board[newx][newy] == 3 || board[newx][newy] == 1) {
				cnt++
			}
		}
		return cnt
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt := count(i, j)
			if board[i][j] == 1 && (cnt < 2 || cnt > 3) {
				//活细胞死亡
				board[i][j] = 3
			} else if board[i][j] == 0 && cnt == 3 {
				//死细胞存活
				board[i][j] = 2
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}
