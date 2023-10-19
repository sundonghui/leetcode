package tuplesameproduct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tupleSameProduct(t *testing.T) {
	//nums := []int{2,3,4,6}
	//assert.Equal(t, 8, tupleSameProduct(nums))
	//nums2 := []int{1,2,4,5,10}
	//assert.Equal(t, 16, tupleSameProduct(nums2))
	nums3 := []int{2, 3, 4, 6, 8, 12}
	assert.Equal(t, 40, tupleSameProduct(nums3))
}
