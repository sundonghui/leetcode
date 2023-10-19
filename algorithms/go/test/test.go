package test

import (
	"fmt"
)

// 幂律科技
// 已知数组 A, B, 如果 A 中元素在 B 数组存在，打印出这个元素的下标，B 数组是不重复的.
// Input:
// A: [5, 3, 1, 5, 4]
// B: [5, 3]
// Output: [0, 1, 3]

func main0() {
	a := []int{5, 3, 1, 5, 4}
	b := []int{5, 3}
	//   fmt.Println(a)
	//   fmt.Println(b)
	bm := make(map[int]struct{})
	for _, v := range b {
		bm[v] = struct{}{}
	}
	output := make([]int, 0, len(a))
	for i, v := range a {
		if _, ok := bm[v]; ok {
			output = append(output, i)
		}
	}
	fmt.Println(output)
	return
}
