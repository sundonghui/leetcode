package insert

func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	ans := [][]int{}
	flag := false
	for _, interval := range intervals {
		if interval[1] < left {
			ans = append(ans, interval)
		} else if interval[0] > right {
			if !flag {
				ans = append(ans, []int{left, right})
				flag = true
			}
			ans = append(ans, interval)
		} else {
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !flag {
		ans = append(ans, []int{left, right})
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
