package constructarr

func constructArr(a []int) []int {
	if len(a) < 1 {
		return []int{}
	}
	length := len(a)
	answer := make([]int, length)

	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = a[i-1] * answer[i-1]
	}

	right := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * right
		right *= a[i]
	}
	return answer
}
