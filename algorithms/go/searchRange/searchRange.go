package searchrange

func searchRange(nums []int, target int) []int {
	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			left, right = mid, mid
			for {
				if (left-1 >= 0) && nums[left-1] == target {
					left--
				} else {
					break
				}
			}
			for {
				if right+1 < len(nums) && nums[right+1] == target {
					right++
				} else {
					break
				}
			}
			return []int{left, right}
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{-1, -1}
}
