package numways

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumWays(t *testing.T) {
	assert.Equal(t, 21, numWays(7))
	//assert.Equal(t, 3, numWays(3))
}
