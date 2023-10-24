package sortthestudents

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortTheStudents(t *testing.T) {
	score := [][]int{
		{10, 6, 9, 1},
		{7, 5, 11, 2},
		{4, 8, 3, 15},
	}
	k := 2
	expected := [][]int{
		{7, 5, 11, 2},
		{10, 6, 9, 1},
		{4, 8, 3, 15},
	}
	assert.Equal(t, expected, sortTheStudents(score, k))
}
