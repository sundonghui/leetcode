package jump

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	assert.Equal(t, 2, jump(nums))
}
