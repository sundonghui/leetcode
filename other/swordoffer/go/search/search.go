package search

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == len(nums) || nums[left] != target {
		return 0
	}
	res := 0
	for i := left; i < len(nums); i++ {
		if nums[i] == target {
			res++
		} else {
			break
		}
	}
	return res
}
