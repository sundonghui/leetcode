package getcommon

func getCommon1(nums1 []int, nums2 []int) int {
	m := make(map[int]struct{}, len(nums1))
	for _, v := range nums1 {
		m[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			return v
		}
	}
	return -1
}

func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}
		if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return -1
}
