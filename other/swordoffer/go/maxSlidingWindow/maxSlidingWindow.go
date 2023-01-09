package maxslidingwindow

import (
	"container/heap"
	"sort"
)

var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

func maxSlidingWindowOverTime(nums []int, k int) []int {
	maxValue := func(list []int) int {
		if len(list) > 0 {
			max := list[0]
			for _, v := range list {
				if v > max {
					max = v
				}
			}
			return max
		}
		return 0
	}
	if k >= len(nums) {
		return []int{maxValue(nums)}
	}
	res := make([]int, 0, len(nums)-k+1)
	i, j := 0, k
	for j <= len(nums) {
		tmp := nums[i:j]
		res = append(res, maxValue(tmp))
		i++
		j++
	}
	return res
}
