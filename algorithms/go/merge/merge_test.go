package merge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Merge(t *testing.T) {
	nums1 := []int{2, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{1, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	expected := []int{1, 2, 2, 3, 5, 6}
	assert.Equal(t, expected, nums1)
}

func Test_Merge2(t *testing.T) {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	m := 3
	nums2 := []int{1, 2, 3}
	n := 3
	merge(nums1, m, nums2, n)
	expected := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, expected, nums1)
}
