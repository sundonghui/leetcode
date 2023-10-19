package countfairpairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countFairPairs(t *testing.T) {
	nums := []int{0, 1, 7, 4, 4, 5}
	lower := 3
	upper := 6
	assert.Equal(t, int64(6), countFairPairs(nums, lower, upper))
}
