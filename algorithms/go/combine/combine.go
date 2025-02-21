package combine

func combine(n int, k int) [][]int {
	var ans [][]int
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if len(temp)+n-cur+1 < k {
			return
		}
		if len(temp) == k {
			value := make([]int, k)
			copy(value, temp)
			ans = append(ans, value)
			return
		}
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		dfs(cur + 1)
	}
	dfs(1)
	return ans
}
