package generate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate(t *testing.T) {
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}
	assert.Equal(t, expected, generate(5))
}
