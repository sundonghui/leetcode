package permute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permute_example1(t *testing.T) {
	//arrange
	input := []int{1, 2, 3}
	expected := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	//act
	result := permute(input)

	//assert
	assert.Equal(t, expected, result)
}

func Test_permute_example2(t *testing.T) {
	//arrange
	input := []int{0, 1}
	expected := [][]int{
		{0, 1},
		{1, 0},
	}

	//act
	result := permute(input)

	//assert
	assert.Equal(t, expected, result)
}
