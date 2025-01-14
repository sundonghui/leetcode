package findminarrowshots

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestFindMinArrowShots(t *testing.T) {
	// Test case 1
	points1 := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	expected1 := 2
	result1 := findMinArrowShots(points1)
	assert.Equal(t, expected1, result1)
}

func TestFindMinArrowShots2(t *testing.T) {
	// Test case 2
	points2 := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	expected2 := 4
	result2 := findMinArrowShots(points2)
	assert.Equal(t, expected2, result2)
}

func TestFindMinArrowShots3(t *testing.T) {
	// Test case 3
	points3 := [][]int{{1, 2}}
	expected3 := 1
	result3 := findMinArrowShots(points3)
	assert.Equal(t, expected3, result3)
}