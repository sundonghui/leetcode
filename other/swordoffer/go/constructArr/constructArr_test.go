package constructarr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructArr(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	expected := []int{120, 60, 40, 30, 24}
	assert.Equal(t, expected, constructArr(a))
}
