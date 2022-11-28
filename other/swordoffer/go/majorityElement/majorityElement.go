package majorityelement

func majorityElement(nums []int) int {
	numsMap := make(map[int]int)
	half := len(nums) / 2
	for _, v := range nums {
		numsMap[v]++
		if numsMap[v] > half {
			return v
		}
	}
	return -1 // 不存在
}
