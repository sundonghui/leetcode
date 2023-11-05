package intersection

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	res := make([]int, 0, len(nums2))
	m := make(map[int]bool, len(nums1))
	for _, v := range nums2 {
		target := v
		left, right := 0, len(nums1)-1
		for left <= right {
			mid := (left + right) / 2
			if nums1[mid] == target && !m[target] {
				m[target] = true
				res = append(res, target)
			}
			if nums1[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return res
}
