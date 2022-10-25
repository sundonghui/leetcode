package findrepeatnumber

func findRepeatNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		// Go map for loop is disorder
		if m[num] > 0 {
			return num
		}
		m[num]++
	}
	return -1
}
