package findduplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDuplicate(t *testing.T) {
	arr := []int{3, 2, 4, 3, 1}
	assert.Equal(t, 3, findDuplicate(arr))
}
