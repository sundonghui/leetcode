package removeelement

func removeElement(nums []int, val int) int {
	for i, v := range nums {
		if v == val {
			return removeElement(append(nums[0:i], nums[i+1:]...), val)
		}
	}
	return len(nums)
}
