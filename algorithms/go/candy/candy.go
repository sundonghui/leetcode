package candy

func candy(ratings []int) int {
	ans := 0
	left := make([]int, len(ratings))
	for i, v := range ratings {
		if i > 0 && v > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 1
	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(right, left[i])
	}
	return ans
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
