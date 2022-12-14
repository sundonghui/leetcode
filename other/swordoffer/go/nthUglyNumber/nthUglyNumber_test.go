package nthuglynumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthUglyNumber(t *testing.T) {
	assert.Equal(t, 12, nthUglyNumber(10))
}
