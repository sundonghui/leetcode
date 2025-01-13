package summaryranges

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	res := []string{}
	for i, n := 0, len(nums); i < n; {
		left := i
		for i++; i < n && nums[i-1] == nums[i]-1; i++ {
		}
		s := strconv.Itoa(nums[left])
		if left < i-1 {
			s += ("->" + strconv.Itoa(nums[i-1]))
		}
		res = append(res, s)
	}
	return res
}
