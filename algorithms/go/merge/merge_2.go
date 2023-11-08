package merge

import (
	"sort"
)

func merge2(intervals [][]int) [][]int {
	list := make([]int, 20)
	for _, interval := range intervals {
		for i := interval[0]; i < interval[1]; i++ {
			list[i] = 1
		}
	}

	res := make([][]int, 0, len(intervals))
	start := 0
	if list[0] == 1 {
		j := 0
		for list[j] == 1 {
			j++
		}
		res = append(res, []int{0, j})
		start = j - 1
	}
	var min, max int
	for i := start; i < len(list); i++ {
		if list[i] == 1 && list[i-1] == 0 {
			min = i
		}
		if i > 0 && list[i] == 0 && list[i-1] == 1 && min > 0 {
			max = i
		}
		if min > 0 && max > 0 {
			res = append(res, []int{min, max})
			min = 0
			max = 0
		}
	}
	return res
}

func mergeSort(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := make([][]int, 0, len(intervals))
	res = append(res, intervals[0])
	for _, interval := range intervals[1:] {
		if res[len(res)-1][1] < interval[0] {
			res = append(res, interval)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], interval[1])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
