package insert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insert_example(t *testing.T) {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	// Output: [[1,5],[6,9]]
	expected := [][]int{{1, 5}, {6, 9}}
	assert.Equal(t, expected, insert(intervals, newInterval))
}

func Test_insert_example2(t *testing.T) {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}
	// Output: [[1,2],[3,10],[12,16]]
	expected := [][]int{{1, 2}, {3, 10}, {12, 16}}
	assert.Equal(t, expected, insert(intervals, newInterval))
}

func Test_insert_example3(t *testing.T) {
	intervals := [][]int{}
	newInterval := []int{5, 7}
	expected := [][]int{{5, 7}}
	assert.Equal(t, expected, insert(intervals, newInterval))
}

func Test_insert_example4(t *testing.T) {
	intervals := [][]int{{1, 5}}
	newInterval := []int{2, 7}
	expected := [][]int{{1, 7}}
	assert.Equal(t, expected, insert(intervals, newInterval))
}
