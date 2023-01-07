package singlenumber

func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
