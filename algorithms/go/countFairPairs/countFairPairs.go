package countfairpairs

import (
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) (ans int64) {
	// 排序不会影响数对的个数，为了能够二分，可以先排序。
	sort.Ints(nums)
	for j, x := range nums {
		// sort.SearchInts 返回元素在切片中的位置，如果元素不存在，返回保持有序性不变，元素插入切片的位置
		r := sort.SearchInts(nums[:j], upper-x+1) 
		l := sort.SearchInts(nums[:j], lower-x)
		ans += int64(r - l)
	}
	return 
}

// 超出时间限制
func countFairPairs1(nums []int, lower int, upper int) int64 {
	cnt := 0
	for i, v := range nums {
		for j := i+1; j <= len(nums)-1; j++ {
			tmp := v + nums[j]
			if lower <= tmp && tmp <= upper {
				cnt++
			}
		}
	}
	return int64(cnt)
}