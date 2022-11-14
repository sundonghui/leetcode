package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchange(t *testing.T) {
	nums := []int{1,2,3,4,5,6}
	expected := []int{1,3,5,6,4,2}
	assert.Equal(t, expected, exchange(nums))	
}