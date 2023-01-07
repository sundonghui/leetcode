package singlenumbers

func singleNumbers(nums []int) []int {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}

	res := make([]int, 0, len(m))
	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
