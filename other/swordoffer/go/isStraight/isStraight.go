package isstraight

import (
	"sort"
)

func isStraight(nums []int) bool {
	list := make([]int, 0, len(nums))
	cnt := 0
	for _, v := range nums {
		if v == 0 {
			cnt++
			continue
		}
		list = append(list, v)
	}
	sort.Ints(list)
	pre := list[0]
	for i, v := range list {
		if i > 0 {
			tmp := v - pre
			if tmp == 1 {
				pre = v
				continue
			}
			if tmp > 1 {
				if tmp-1 <= cnt {
					cnt -= (tmp - 1)
					pre = v
					continue
				}
			}
			return false
		}
	}
	return true
}
