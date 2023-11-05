package intersection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intersection(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	assert.Equal(t, []int{2}, intersection(nums1, nums2))
}

func Test_intersection1(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	assert.Equal(t, []int{4, 9}, intersection(nums1, nums2))
}
