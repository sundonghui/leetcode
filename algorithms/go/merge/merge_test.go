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

func Test_Merge2_1(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
	assert.Equal(t, expected, merge2(intervals))
}

func Test_Merge2_2(t *testing.T) {
	intervals := [][]int{{0, 3}, {2, 6}, {8, 10}, {15, 18}}
	expected := [][]int{{0, 6}, {8, 10}, {15, 18}}
	assert.Equal(t, expected, mergeSort(intervals))
}

// [[1,4],[5,6]]
func Test_Merge2_3(t *testing.T) {
	intervals := [][]int{{1, 4}, {5, 6}, {0, 0}}
	expected := [][]int{{0, 0}, {1, 4}, {5, 6}}
	assert.Equal(t, expected, mergeSort(intervals))
}
