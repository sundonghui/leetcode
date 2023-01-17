package maxprofit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	list := []int{7, 1, 5, 3, 6, 4}
	assert.Equal(t, 5, maxProfit(list))
}
