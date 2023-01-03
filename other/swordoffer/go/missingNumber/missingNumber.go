package missingnumber

func missingNumber(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	m := make(map[int]bool, len(nums))
	for _, n := range nums {
		m[n] = true
	}
	for i := 0; ; i++ {
		if !m[i] {
			return i
		}
	}
}
