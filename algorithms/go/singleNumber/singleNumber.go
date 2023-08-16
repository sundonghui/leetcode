package singlenumber

func singleNumber(nums []int) int {
	// 使用 hash 统计每一个数字出现的概率
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	// 找出 hash 中出现次数为 1 的数字
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return nums[0]
}

func singleNumber2(nums []int) int {
	// 使用异或运算
	// 知识点：
	// 1. 0 和任何数字进行异或运算，结果仍是原来的数字
	// 2. 两个相同的数字进行异或运算，结果为 0
	// 3. 异或运算满足交换律和结合律
	r := 0
	for _, v := range nums {
		r ^= v
	}
	return r
}
