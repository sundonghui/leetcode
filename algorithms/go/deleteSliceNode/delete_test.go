package deleteslicenode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Delete(t *testing.T) {
	expected := []int{1, 2, 3}
	arr := []int{1, 2, 3, 4}
	assert.Equal(t, expected, Delete(arr, 3))
}

func Test_Delete_empty(t *testing.T) {
	expected := []int{}
	arr := []int{1}
	assert.Equal(t, expected, Delete(arr, 0))
}
