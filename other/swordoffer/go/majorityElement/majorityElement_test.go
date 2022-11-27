package majorityelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	expected := 2
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	assert.Equal(t, expected, majorityElement(nums))
}