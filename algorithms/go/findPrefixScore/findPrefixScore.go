package findprefixscore

import "fmt"

func findPrefixScore(nums []int) []int64 {
	max := nums[0]
	nums1 := make([]int, len(nums))
	for i, v := range nums {
		if v > max {
			max = v
		}
		nums1[i] = v + max
	}
	fmt.Println(nums1)

	nums2 := make([]int64, len(nums1))
	for i, v := range nums1 {
		if i == 0 {
			nums2[i] = int64(v)
		} else {
			nums2[i] = int64(v) + nums2[i-1]
		}
	}
	return nums2
}
