package findduplicate

import "sort"

func findDuplicate2(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

func findDuplicate(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return ans
}

func findDuplicate3(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
