package reversepairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversePairs(t *testing.T) {
	nums := []int{7, 5, 6, 4}

	assert.Equal(t, 5, reversePairs(nums))
}
