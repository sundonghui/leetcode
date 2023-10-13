package findthearrayconcval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheArrayConcVal(t *testing.T) {
	nums := []int{7, 52, 2, 4}
	assert.Equal(t, int64(596), findTheArrayConcVal(nums))
}
