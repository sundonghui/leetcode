package exchange

func exchange(nums []int) []int {
	r := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for _, v := range nums {
		if v%2 == 1 {
			r[left] = v
			left++
		} else {
			r[right] = v
			right--
		}
	}
	return r
}