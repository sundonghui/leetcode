package sortthestudents

import "sort"

func sortTheStudents(score [][]int, k int) [][]int {
	m := make(map[int][]int, len(score))
	list := make([]int, 0, len(score))
	for _, v := range score {
		tmp := v[k]
		m[tmp] = v
		list = append(list, tmp)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j]
	})
	result := make([][]int, 0, len(score))
	for _, v := range list {
		result = append(result, m[v])
	}
	return result
}
