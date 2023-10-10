package sumdistance

import (
	"sort"
)

func sumDistance(nums []int, s string, d int) int {
	const mod = 1e9 + 7
	n := len(nums)
	pos := make([]int, n)
	for i, ch := range s {
		if ch == 'R' {
			pos[i] = nums[i] + d
		} else {
			pos[i] = nums[i] - d
		}
	}
	sort.Ints(pos)
	res := 0
	for i := 1; i < n; i++ {
		res += (pos[i] - pos[i-1]) % mod * i * (n - i) % mod
		res %= mod
	}
	return res
}
