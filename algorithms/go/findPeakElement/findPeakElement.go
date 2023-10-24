package findpeakelement

import "math"

func findPeakElement(nums []int) int {
	get := func(i int) int {
		if i < 0 || i >= len(nums) {
			return math.MinInt64
		}
		return nums[i]
	}

	left, right := 0, len(nums)-1
	for {
		mid := (left + right) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
