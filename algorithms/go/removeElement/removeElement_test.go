package removeelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{1, 2, 3, 4, 4, 4, 4}
	assert.Equal(t, 3, removeElement(nums, 4))
}
