package getcommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getCommon(t *testing.T) {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4}
	assert.Equal(t, 2, getCommon(nums1, nums2))
}

func Test_getCommon1(t *testing.T) {
	nums1 := []int{1, 2, 3, 6}
	nums2 := []int{2, 3, 4, 5}
	assert.Equal(t, 2, getCommon(nums1, nums2))
}
