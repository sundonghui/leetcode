package pivotindex

func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	sum := 0
	for i, v := range nums {
		if 2*sum == total-v {
			return i
		}
		sum += v
	}

	return -1
}
