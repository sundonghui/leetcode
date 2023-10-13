package searchinsert

func searchInsert(nums []int, target int) int {
	n := len(nums)
	start, end := 0, n-1
	ans := n
	for start <= end {
		mid := start + (end-start)>>1
		if target <= nums[mid] {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}
