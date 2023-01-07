package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for _, n := range nums {
		r := target - n
		if _, ok := m[r]; ok {
			return []int{r, n}
		}
		m[n]++
	}
	return []int{}
}
