package getleastnumbers

func getLeastNumbers(arr []int, k int) []int {
	nums := quickSort(arr, 0, len(arr)-1)
	return nums[:k]
}

func quickSort(nums []int, left, right int) []int {
	if left < right {
		index := partition(nums, left, right)
		quickSort(nums, left, index-1)
		quickSort(nums, index+1, right)
	}
	return nums
}

func partition(nums []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if nums[i] < nums[pivot] {
			swap(nums,i, index)
			index++
		}
	}
	swap(nums, pivot, index-1)
	return index-1
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}