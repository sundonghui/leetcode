package merge

func merge(nums1 []int, m int, nums2 []int, n int) {
	for m+n > 0 {
		if m <= 0 && n > 0 {
			nums1[n-1] = nums2[n-1]
			n--
			continue
		}
		if n == 0 {
			break
		}
		a := nums1[m-1]
		b := nums2[n-1]
		if a > b {
			nums1[m+n-1] = a
			m--
		} else {
			nums1[m+n-1] = b
			n--
		}
	}
}
