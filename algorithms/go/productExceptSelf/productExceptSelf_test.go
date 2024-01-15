package productexceptself

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_productExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}
	assert.Equal(t, expected, productExceptSelf(nums))
}
