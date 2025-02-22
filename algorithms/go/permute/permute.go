package permute

func permute(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	path := make([]int, n)
	used := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		if x == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < n; i++ {
			if !used[i] {
				path[x] = nums[i]
				used[i] = true
				dfs(x + 1)
				used[i] = false
			}
		}
	}
	dfs(0)
	return ans
}
