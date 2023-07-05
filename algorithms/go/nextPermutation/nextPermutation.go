package nextpermutation

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 找到 nums[i] < nums[j]
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}

		nums[i], nums[k] = nums[k], nums[i]
	}

	// 将 i - j 倒序排列
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
