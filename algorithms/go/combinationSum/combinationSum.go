package combinationsum

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0, 150)
	comb := []int{}
	var dfs func(target int, idx int)
	dfs = func(target int, idx int) {
		if len(ans) == 150 {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int{}, comb...))
			return
		}
		if idx == len(candidates) {
			return
		}
		if target >= candidates[idx] {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
		dfs(target, idx+1)
	}
	dfs(target, 0)
	return ans
}
