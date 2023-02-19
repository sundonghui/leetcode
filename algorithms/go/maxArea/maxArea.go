package maxarea

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		long := j - i
		high := height[i]
		if height[i] > height[j] {
			high = height[j]
			j--
		} else {
			high = height[i]
			i++
		}
		if long*high > max {
			max = long * high
		}
	}
	return max
}
