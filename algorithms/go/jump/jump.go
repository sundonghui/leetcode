package jump

func jump(nums []int) int {
	step := 0
	end := 0
	maxPosition := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			step++
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
