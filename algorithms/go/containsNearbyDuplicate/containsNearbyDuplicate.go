package containsnearbyduplicate

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if _, ok := m[num]; ok && abs(m[num], i) <= k {
			return true
		}
		m[num] = i
	}
	return false
}

func abs(i, j int) int {
	if i-j < 0 {
		return j - i
	}
	return i - j
}
