package maxkelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxKelements(t *testing.T) {
	nums := []int{1, 10, 3, 3, 3}
	k := 3
	assert.Equal(t, int64(17), maxKelements(nums, k))
}
