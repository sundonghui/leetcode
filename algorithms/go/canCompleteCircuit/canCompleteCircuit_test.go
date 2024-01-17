package cancompletecircuit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanCompleteCircuit(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	assert.Equal(t, 3, canCompleteCircuit(gas, cost))
}

func Test_CanNotCompleteCircuit(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	assert.Equal(t, -1, canCompleteCircuit(gas, cost))
}
