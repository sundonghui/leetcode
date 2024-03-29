package splitarray

func splitArray(nums []int, m int) int {
	left, right := 0, 0
	for i := 0; i < len(nums); i++ {
		right += nums[i]
		if left < nums[i] {
			left = nums[i]
		}
	}
	for left < right {
		mid := (right-left)/2 + left
		if check(nums, mid, m) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(nums []int, x, m int) bool {
	sum, cnt := 0, 1
	for i := 0; i < len(nums); i++ {
		if sum+nums[i] > x {
			cnt++
			sum = nums[i]
		} else {
			sum += nums[i]
		}
	}
	return cnt <= m
}
