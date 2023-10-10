package splitnum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitNum(t *testing.T) {
	assert.Equal(t, 59, splitNum(4325))
}
