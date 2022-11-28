package getleastnumbers

func getLeastNumbers(arr []int, k int) []int {
	nums := quickSort(arr, 0, len(arr)-1)
	return nums[:k]
}

func quickSort(nums []int, left, right int) []int {
	if left > right {
		return nil
	}

	i, j, pivot := left, right, nums[left]
	for i < j {
		// 寻找小于主元的右边元素
		for i < j && nums[j] >= pivot {
			j--
		}
		// 寻找大于主元的左边元素
		for i < j && nums[i] <= pivot {
			i++
		}
		// 交换i/j下标元素
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 交换元素
	nums[i], nums[left] = nums[left], nums[i]
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
	return nums
}
