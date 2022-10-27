package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for index, v := range nums {
		m[target-v] = index
	}

	for index, v := range nums {
		if v, ok := m[v]; ok {
			if v != index {
				return []int{v, index}
			}
		}
	}
	return nil
}
