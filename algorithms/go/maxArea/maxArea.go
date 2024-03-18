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

func maxArea1(height []int) int {
	res := 0
	l := 0
	r := len(height) - 1
	for l < r {
		temp := min(height[l], height[r])
		res = max(res, temp*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
