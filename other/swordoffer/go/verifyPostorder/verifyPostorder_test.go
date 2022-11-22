package verifypostorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyPostorder(t *testing.T) {
	order1 := []int{1, 6, 3, 2, 5}
	order2 := []int{1, 3, 2, 6, 5}
	assert.Equal(t, false, verifyPostorder(order1))
	assert.Equal(t, true, verifyPostorder(order2))
}
