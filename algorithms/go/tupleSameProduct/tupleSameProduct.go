package tuplesameproduct

func tupleSameProduct(nums []int) int {
	m := make(map[int]int, len(nums)*len(nums))
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			m[v*nums[j]]++
		}
	}

	cnt := 0
	for _, v := range m {
		// 对于 N 个相同的值，每两个值组成一组，每组存在 4 种不同的排列组合顺序，即 C(n, 2)
		// cnt(a*b) * (cnt(a*b)-1) / 2 * 8 = cnt(a*b) * (cnt(a*b)-1) * 4
		cnt += v * (v - 1) * 4
	}
	return cnt
}
