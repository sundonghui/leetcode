package removeduplicates

func removeDuplicates(nums []int) int {
	i, j := 0, 1
	if len(nums) < 2 {
		return len(nums)
	}
	for j < len(nums) {
		pre := nums[i]
		next := nums[j]
		if pre == next {
			nums = append(nums[0:i], nums[i+1:]...)
			continue
		}
		i++
		j++
	}
	return len(nums)
}

func RemoveDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i, j := 0, 1
	for j < len(nums) {
		pre := nums[i]
		next := nums[j]
		if pre == next {
			nums = append(nums[0:i], nums[i+1:]...)
			continue
		}
		i++
		j++
	}
	return len(nums)
}

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	slow, fast := 2, 2
	for fast < len(nums) {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
