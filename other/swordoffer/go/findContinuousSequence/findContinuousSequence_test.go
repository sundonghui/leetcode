package findcontinuoussequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindContinuousSequence(t *testing.T) {
	expected := [][]int{
		{2, 3, 4},
		{4, 5},
	}
	assert.Equal(t, expected, findContinuousSequence(9))
}
