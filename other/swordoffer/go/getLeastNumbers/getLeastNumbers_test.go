package getleastnumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLeastNumbers(t *testing.T) {
	arr := []int{3, 2, 1}
	expected := []int{1, 2}
	assert.Equal(t, expected, getLeastNumbers(arr, 2))
}
