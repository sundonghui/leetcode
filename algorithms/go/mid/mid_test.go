package mid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4}
	assert.Equal(t, 2, search(nums, 2))
}
