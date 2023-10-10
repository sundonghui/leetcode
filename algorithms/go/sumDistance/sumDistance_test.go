package sumdistance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumDistance(t *testing.T) {
	nums := []int{-2, 0, 2}
	s := "RLL"
	assert.Equal(t, 8, sumDistance(nums, s, 3))

}
