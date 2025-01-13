package summaryranges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_summaryRanges_exmple1(t *testing.T) {
	//arrange
	nums := []int{0, 1, 2, 4, 5, 7}
	expected := []string{"0->2", "4->5", "7"}

	//act
	actual := summaryRanges(nums)

	//assert
	assert.Equal(t, expected, actual)
}

func Test_summaryRanges_exmple2(t *testing.T) {
	//arrange
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	expected := []string{"0", "2->4", "6", "8->9"}

	//act
	actual := summaryRanges(nums)

	//assert
	assert.Equal(t, expected, actual)
}
