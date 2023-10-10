package splitnum

import (
	"fmt"
	"sort"
)

func splitNum(num int) int {
	arr := []int{}
	for num >= 10 {
		arr = append(arr, num%10)
		num = num / 10
	}
	arr = append(arr, num)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	fmt.Println(arr)
	r := 0
	p := 1
	for i := 0; i < len(arr); i += 2 {
		cur := arr[i]
		next := 0
		if i+1 < len(arr) {
			next = arr[i+1]
		}
		tmp := cur + next
		r += tmp * p
		p *= 10
	}
	return r
}
