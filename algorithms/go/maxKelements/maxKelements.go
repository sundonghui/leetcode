package maxkelements

import (
	"container/heap"
	"math"
	"sort"
)

func maxKelements1(nums []int, k int) int64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	sum := 0
	for k > 0 {
		val := nums[0]
		sum += val
		val = int(math.Ceil(float64(val) / 3))
		nums[0] = val
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] > nums[j]
		})
		k--
	}
	return int64(sum)
}

func maxKelements(nums []int, k int) (ans int64) {
	h := hp{nums}
	heap.Init(&h) // 原地堆化
	for ; k > 0; k-- {
		ans += int64(h.IntSlice[0]) // 堆顶
		h.IntSlice[0] = (h.IntSlice[0] + 2) / 3
		heap.Fix(&h, 0)
	}
	return
}

type hp struct{ sort.IntSlice }

func (h *hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (h *hp) Push(any)           {}
func (h *hp) Pop() (_ any)       { return }
