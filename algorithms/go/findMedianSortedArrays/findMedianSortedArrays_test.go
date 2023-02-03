package findmediansortedarrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	assert.Equal(t, 2.000, findMedianSortedArrays(nums1, nums2))
}
